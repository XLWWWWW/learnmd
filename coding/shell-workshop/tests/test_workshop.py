#!/usr/bin/env python3
"""Shell 黑盒测试。Python 只提供断言、超时与临时目录，不参与习题实现。"""
import argparse
import contextlib
import os
from pathlib import Path
import signal
import subprocess
import sys
import tarfile
import tempfile
import time
import unittest

ROOT = Path(__file__).resolve().parents[1]
PARSER = argparse.ArgumentParser()
PARSER.add_argument('--mode', choices=('practice', 'solution'), default='practice')
PARSER.add_argument('--exercise', default='all')
PARSER.add_argument('--bash', default='/bin/bash')
OPTIONS = PARSER.parse_args()
SCRIPTS = ROOT / ('solutions' if OPTIONS.mode == 'solution' else 'practice')


class WorkshopTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix='shell lab ')
        self.dir = Path(self.temp.name)
        self.processes = []
        self.env = os.environ.copy()
        # 防止调用者交互式 Shell 的启动配置改变题目语义。
        for key in ('BASH_ENV', 'ENV', 'SHELLOPTS', 'BASHOPTS', 'CDPATH',
                    'APP_ENV', 'APP_PORT', 'WORK_DIR', 'GLOBIGNORE'):
            self.env.pop(key, None)
        self.env.update(LC_ALL='C', TMPDIR=str(self.dir))

    def tearDown(self):
        for process in self.processes:
            # 连同未回收的后代一起清理测试自己的进程组。
            with contextlib.suppress(ProcessLookupError):
                os.killpg(process.pid, signal.SIGKILL)
            process.communicate()
        self.temp.cleanup()

    def launch(self, number, args=(), env=None):
        merged = self.env.copy()
        merged.update(env or {})
        process = subprocess.Popen(
            [OPTIONS.bash, str(SCRIPTS / ('e%02d.sh' % number))] + list(map(str, args)),
            cwd=self.dir, env=merged, stdin=subprocess.PIPE, stdout=subprocess.PIPE,
            stderr=subprocess.PIPE, start_new_session=True)
        self.processes.append(process)
        return process

    def finish(self, process, data=b'', timeout=6):
        try:
            out, err = process.communicate(data, timeout=timeout)
        except subprocess.TimeoutExpired:
            os.killpg(process.pid, signal.SIGKILL)
            out, err = process.communicate()
            self.fail('脚本超时（可能存在循环或没有回收子进程）: ' + repr((out, err)))
        return process.returncode, out, err

    def run_case(self, number, args=(), data=b'', env=None):
        return self.finish(self.launch(number, args, env), data)

    def ok(self, number, args=(), output=b'', data=b'', env=None):
        rc, out, err = self.run_case(number, args, data, env)
        self.assertEqual(rc, 0, err.decode(errors='replace'))
        self.assertEqual(out, output)
        return err

    def bad(self, number, args=(), code=2, data=b'', env=None):
        rc, out, err = self.run_case(number, args, data, env)
        self.assertEqual(rc, code, repr((out, err)))
        self.assertEqual(out, b'', '错误输入不应输出成功数据')
        self.assertTrue(err, '错误需要向 stderr 给出提示')

    def file(self, name, data=b''):
        path = self.dir / name
        path.write_bytes(data)
        return path

    def worker(self, code, name='worker with spaces'):
        # 受控的本地测试程序；在被测 Shell 中必须把路径作为一个参数传递。
        path = self.file(name, ('#!' + sys.executable + '\n' + code + '\n').encode())
        path.chmod(0o700)
        return path

    def await_file(self, path, process):
        deadline = time.monotonic() + 4
        while time.monotonic() < deadline:
            if path.exists() and path.stat().st_size:
                return
            if process.poll() is not None:
                self.fail('子任务尚未就绪，脚本已经退出: ' + repr(process.communicate()))
            time.sleep(0.01)
        self.fail('等待子任务就绪超时')

    def test_e01_boundaries(self):
        self.file('expanded.log')
        self.ok(1, ['hello shell', '', '*.log', '-n', '中文'],
                'count=5\n1=<hello shell>\n2=<>\n3=<*.log>\n4=<-n>\n5=<中文>\n'.encode())

    def test_e01_empty_and_newline(self):
        self.ok(1, output=b'count=0\n')
        self.ok(1, ['a\nb'], b'count=1\n1=<a\nb>\n')

    def test_e02_defaults(self):
        self.ok(2, output=b'env=dev\nport=8080\n')
        self.ok(2, output=b'env=dev\nport=8080\n', env={'APP_ENV': '', 'APP_PORT': ''})

    def test_e02_overrides_and_usage(self):
        self.ok(2, output=b'env=pre prod\nport=9090\n', env={'APP_ENV': 'pre prod', 'APP_PORT': '9090'})
        self.bad(2, ['extra'])

    def test_e03_numbers(self):
        for value in ('0', '00', '008', '010', '255'):
            with self.subTest(value=value):
                n = int(value)
                self.ok(3, [value], ('%d\n' % (n * (n + 1) // 2)).encode())

    def test_e03_invalid(self):
        for args in ([], [''], ['256'], ['9999'], ['-1'], ['1+2'], [' 8'], ['1', '2'], ['$(touch nope)']):
            with self.subTest(args=args): self.bad(3, args)

    def test_e04_parse(self):
        self.ok(4, ['--name', 'A B'], b'hello=A B\n')
        self.ok(4, ['--times', '2', '--name', '-bot', '--'], b'hello=-bot\nhello=-bot\n')

    def test_e04_invalid(self):
        for args in ([], ['--name'], ['--name', ''], ['--name', 'x', '--times'],
                     ['--name', 'x', '--times', '0'], ['--name', 'x', '--times', '06'],
                     ['--name', 'x', '--name', 'y'], ['--name', 'x', '--', 'extra'],
                     ['--name', 'x', '--times', '2', '--times', '3'], ['--unknown']):
            with self.subTest(args=args): self.bad(4, args)

    def test_e05_preserves_text(self):
        path = self.file('input file', b'  a  \n\nback\\slash\nlast')
        self.ok(5, [path], b'1:  a  \n2:\n3:back\\slash\n4:last\n')

    def test_e05_empty_and_errors(self):
        self.ok(5, [self.file('empty')])
        self.ok(5, [self.file('-n', b'\n')], b'1:\n')
        self.bad(5, [], 2)
        self.bad(5, [self.dir / 'missing'], 1)
        self.bad(5, [self.dir], 1)

    def test_e06_report(self):
        self.ok(6, [ROOT / 'fixtures/access.log'], b'api\t3\t2\nworker\t2\t0\n')
        self.ok(6, [self.file('empty')])
        self.ok(6, [self.file('last', b't INFO api 200')], b'api\t1\t0\n')

    def test_e06_rejects_partial_report(self):
        for line in (b't INFO api 20', b't INFO api nope', b't DEBUG api 200',
                     b't INFO api 200 extra', b't INFO API 200'):
            with self.subTest(line=line):
                path = self.file('bad log', b't INFO api 200\n' + line)
                self.bad(6, [path], 1)
        self.bad(6, [self.dir / 'missing'], 1)
        self.bad(6, [])

    def test_e06_sort_failure(self):
        fake = self.dir / 'fake'; fake.mkdir()
        self.worker('import sys\nsys.exit(9)', 'fake/sort')
        self.bad(6, [ROOT / 'fixtures/access.log'], 1,
                 env={'PATH': str(fake) + os.pathsep + self.env['PATH']})

    def test_e07_unusual_filenames_and_symlinks(self):
        folder = self.dir / 'logs'; folder.mkdir()
        expected = []
        for name in ('a b.log', 'a\nb.log', '-n.log', '.hidden.log', '*.log'):
            path = folder / name; path.write_bytes(b''); expected.append(os.fsencode(path))
        (folder / 'x.txt').touch()
        (folder / 'directory.log').mkdir()
        outside = self.file('outside.log')
        (folder / 'link.log').symlink_to(outside)
        (folder / 'linked-dir').symlink_to(self.dir, target_is_directory=True)
        rc, out, err = self.run_case(7, [folder])
        self.assertEqual(rc, 0, repr(err))
        self.assertTrue(out.endswith(b'\0'))
        self.assertEqual(sorted(out[:-1].split(b'\0')), sorted(expected))

    def test_e07_relative_dash_and_empty(self):
        folder = self.dir / '-logs'; folder.mkdir()
        self.ok(7, ['-logs'])
        (folder / 'x.log').touch()
        self.ok(7, ['-logs'], b'./-logs/x.log\0')
        self.bad(7, [self.dir / 'missing'], 1)
        link = self.dir / 'alias'; link.symlink_to(folder, target_is_directory=True)
        self.bad(7, [link], 1)

    def test_e07_symlink_with_trailing_slash(self):
        folder = self.dir / 'logs'; folder.mkdir()
        link = self.dir / 'alias'; link.symlink_to(folder, target_is_directory=True)
        self.bad(7, [str(link) + '/'], 1)
        self.ok(7, [str(folder) + '/'])

    def test_e08_exec_preserves_arguments_and_status(self):
        self.ok(8, ['--run', '--', OPTIONS.bash, '-c', 'printf "<%s>\\n" "$@"', '_', 'a b', '', '*'],
                b'<a b>\n<>\n<*>\n')
        rc, _, _ = self.run_case(8, ['--run', '--', OPTIONS.bash, '-c', 'exit 7'])
        self.assertEqual(rc, 7)

    def test_e08_dry_run(self):
        self.ok(8, ['--dry-run', '--', 'touch', 'not-created'], b'touch\nnot-created\n')
        self.assertFalse((self.dir / 'not-created').exists())
        self.ok(8, ['--dry-run', '--', 'printf', 'a b', '*', ''], b"printf\na\\ b\n\\*\n''\n")
        self.bad(8, ['--run', '--'])
        self.bad(8, ['wat', '--', 'true'])

    def test_e09_literal_config(self):
        self.ok(9, [ROOT / 'fixtures/app.conf'], b'name=hello shell\nport=8080\n')
        path = self.file('literal.conf', b'APP_NAME=$(touch pwned)=literal\nPORT=00080')
        self.ok(9, [path], b'name=$(touch pwned)=literal\nport=80\n')
        self.assertFalse((self.dir / 'pwned').exists())

    def test_e09_invalid_configs(self):
        for content in (b'PORT=80', b'APP_NAME=x\nPORT=0', b'APP_NAME=x\nPORT=65536',
                        b'APP_NAME=x\nPORT=80\nPORT=90', b'APP_NAME=x\nPORT=80\nOTHER=x',
                        b'APP_NAME=\nPORT=80', b'APP_NAME=x\r\nPORT=80\r\n',
                        b'APP_NAME=x\nPORT=1+2', b'APP_NAME=x\n PORT=80'):
            with self.subTest(content=content): self.bad(9, [self.file('bad.conf', content)])
        self.bad(9, [self.dir / 'missing'], 1)

    def test_e10_binary_and_replace(self):
        path = self.file('output file', b'old')
        self.ok(10, [path], data=b'a\0b\nlast')
        self.assertEqual(path.read_bytes(), b'a\0b\nlast')
        self.assertEqual(path.stat().st_mode & 0o777, 0o600)
        self.assertEqual(list(self.dir.glob('.shell-write.*')), [])
        self.ok(10, ['-result'], data=b'')
        self.assertEqual((self.dir / '-result').read_bytes(), b'')

    def test_e10_invalid_destinations(self):
        original = self.file('original', b'old')
        link = self.dir / 'link'; link.symlink_to(original)
        self.bad(10, [link], 1, data=b'new')
        self.assertEqual(original.read_bytes(), b'old')
        self.bad(10, [self.dir], 1)
        self.bad(10, [self.dir / 'missing/file'], 1)
        self.bad(10, [])

    def test_e10_failed_writer_preserves_original(self):
        fake = self.dir / 'fake'; fake.mkdir()
        self.worker('import sys\nsys.stdout.write("partial")\nsys.exit(9)', 'fake/cat')
        target = self.file('target', b'old')
        rc, _, _ = self.run_case(10, [target], env={'PATH': str(fake) + os.pathsep + self.env['PATH']})
        self.assertNotEqual(rc, 0)
        self.assertEqual(target.read_bytes(), b'old')
        self.assertEqual(list(self.dir.glob('.shell-write.*')), [])

    def test_e11_backup_restores_contents(self):
        source = self.dir / 'source'; source.mkdir()
        (source / 'a b.txt').write_bytes(b'hello\0world')
        (source / '.hidden').write_bytes(b'hidden')
        (source / 'sub').mkdir()
        archive = self.dir / 'backup.tar.gz'
        self.ok(11, [source, archive])
        with tarfile.open(archive) as reader:
            self.assertEqual(reader.extractfile('./a b.txt').read(), b'hello\0world')
            self.assertEqual(reader.extractfile('./.hidden').read(), b'hidden')
            self.assertTrue(reader.getmember('./sub').isdir())
        self.assertEqual(list(self.dir.glob('.shell-backup.*')), [])

    def test_e11_refuses_overwrite_and_nested_archive(self):
        source = self.dir / 'source'; source.mkdir()
        archive = self.file('keep.tar.gz', b'keep')
        self.bad(11, [source, archive], 1)
        self.assertEqual(archive.read_bytes(), b'keep')
        self.bad(11, [source, source / 'nested.tar.gz'], 2)
        alias = self.dir / 'alias'; alias.symlink_to(source, target_is_directory=True)
        self.bad(11, [source, alias / 'nested.tar.gz'], 2)
        self.bad(11, ['source', archive], 2)

    def test_e11_tar_failure_no_archive(self):
        source = self.dir / 'source'; source.mkdir()
        fake = self.dir / 'fake'; fake.mkdir()
        self.worker('import sys\nsys.exit(9)', 'fake/tar')
        archive = self.dir / 'backup.tar.gz'
        rc, _, _ = self.run_case(11, [source, archive], env={'PATH': str(fake) + os.pathsep + self.env['PATH']})
        self.assertNotEqual(rc, 0)
        self.assertFalse(archive.exists())
        self.assertEqual(list(self.dir.glob('.shell-backup.*')), [])

    def test_e11_source_symlink_with_trailing_slash(self):
        source = self.dir / 'source'; source.mkdir()
        link = self.dir / 'alias'; link.symlink_to(source, target_is_directory=True)
        self.bad(11, [str(link) + '/', self.dir / 'archive.tar.gz'], 1)
        self.ok(11, [str(source) + '/', self.dir / 'archive.tar.gz'])

    def test_e12_eventual_success(self):
        worker = self.worker('import pathlib, sys\np=pathlib.Path(sys.argv[1])\nn=int(p.read_text())+1 if p.exists() else 1\np.write_text(str(n))\nsys.exit(0 if n==3 else 7)')
        count = self.dir / 'count'
        self.ok(12, ['5', '0', '--', worker, count])
        self.assertEqual(count.read_text(), '3')

    def test_e12_exhaustion_and_usage(self):
        worker = self.worker('import pathlib, sys\np=pathlib.Path(sys.argv[1])\np.write_text(p.read_text()+"x" if p.exists() else "x")\nsys.exit(7)')
        count = self.dir / 'count'
        rc, _, _ = self.run_case(12, ['2', '0', '--', worker, count])
        self.assertEqual(rc, 7)
        self.assertEqual(count.read_text(), 'xx')
        self.bad(12, ['0', '0', '--', worker, count])
        self.bad(12, ['2', '0', '--'])

    def test_e13_cleanup_on_success_and_failure(self):
        worker = self.worker('import os, pathlib, sys\np=pathlib.Path(os.environ["WORK_DIR"])\nassert p.is_dir()\n(p/"data").write_text("x")\nprint(p)\nsys.exit(int(sys.argv[1]))')
        for code in (0, 7):
            with self.subTest(code=code):
                rc, out, _ = self.run_case(13, ['--', worker, str(code)])
                self.assertEqual(rc, code)
                self.assertFalse(Path(os.fsdecode(out.strip())).exists())

    def test_e13_preserves_stdin(self):
        worker = self.worker('import sys\nsys.stdout.buffer.write(sys.stdin.buffer.read())')
        self.ok(13, ['--', worker], b'input\0data', data=b'input\0data')

    def test_e13_term_cleans_and_waits(self):
        ready = self.dir / 'ready'
        worker = self.worker('import os, pathlib, signal, sys\npathlib.Path(sys.argv[1]).write_text(os.environ["WORK_DIR"]+"\\n"+str(os.getpid()))\nsignal.pause()')
        process = self.launch(13, ['--', worker, ready])
        self.await_file(ready, process)
        workspace, pid = ready.read_text().splitlines()
        self.assertTrue(Path(workspace).is_dir())
        process.send_signal(signal.SIGTERM)
        rc, _, _ = self.finish(process)
        self.assertEqual(rc, 143)
        self.assertFalse(Path(workspace).exists())
        with self.assertRaises(ProcessLookupError): os.kill(int(pid), 0)

    def test_e14_contention_and_release(self):
        lock = self.dir / 'job.lock'; ready = self.dir / 'ready'
        worker = self.worker('import pathlib, signal, sys\npathlib.Path(sys.argv[1]).write_text("ready")\nsignal.pause()')
        process = self.launch(14, [lock, '--', worker, ready])
        self.await_file(ready, process)
        self.bad(14, [lock, '--', 'touch', str(self.dir / 'should-not-run')], 75)
        self.assertFalse((self.dir / 'should-not-run').exists())
        self.assertTrue(lock.is_dir())
        process.send_signal(signal.SIGTERM)
        rc, _, _ = self.finish(process)
        self.assertEqual(rc, 143)
        self.assertFalse(lock.exists())
        self.ok(14, [lock, '--', OPTIONS.bash, '-c', 'exit 0'])
        self.assertFalse(lock.exists())

    def test_e14_failure_releases_own_lock(self):
        lock = self.dir / 'job.lock'
        rc, _, _ = self.run_case(14, [lock, '--', OPTIONS.bash, '-c', 'exit 9'])
        self.assertEqual(rc, 9)
        self.assertFalse(lock.exists())
        lock.mkdir(); self.file('job.lock/owner', b'other')
        self.bad(14, [lock, '--', 'true'], 75)
        self.assertEqual((lock / 'owner').read_bytes(), b'other')

    def test_e15_bounded_concurrency_and_all_results(self):
        # flock 仅是 Python 测试计数器的互斥；答案脚本无需 Linux flock 命令。
        worker = self.worker('''import fcntl, json, os, pathlib, sys, time
p=pathlib.Path(os.environ["STATE"])
def update(delta):
    with p.open("r+") as f:
        fcntl.flock(f, fcntl.LOCK_EX)
        s=json.load(f); s["active"]+=delta
        s["max"]=max(s["max"], s["active"])
        if delta<0: s["items"].append(sys.argv[1])
        f.seek(0); json.dump(s,f); f.truncate()
update(1)
time.sleep(0.12)
update(-1)
sys.exit(7 if sys.argv[1]=="bad" else 0)''')
        import json
        items = ['a b', '', 'bad', '*', '-n', 'last']
        for limit in (1, 2, 4):
            with self.subTest(limit=limit):
                state = self.file('state', b'{"active":0,"max":0,"items":[]}')
                rc, _, err = self.run_case(15, [str(limit), worker, '--'] + items, env={'STATE': str(state)})
                self.assertEqual(rc, 1, repr(err))
                data = json.loads(state.read_text())
                self.assertEqual(data['active'], 0)
                self.assertLessEqual(data['max'], limit)
                self.assertCountEqual(data['items'], items)

    def test_e15_actually_runs_in_parallel(self):
        # 两个任务必须同时到达屏障，串行实现会超时失败；不凭跑得快慢判断。
        worker = self.worker('''import os, pathlib, sys, time
base=pathlib.Path(os.environ["BARRIER"])
(base/sys.argv[1]).touch()
deadline=time.monotonic()+3
while len(list(base.iterdir()))<2:
    if time.monotonic()>deadline: sys.exit(8)
    time.sleep(0.01)''')
        barrier = self.dir / 'barrier'; barrier.mkdir()
        self.ok(15, ['2', worker, '--', 'one', 'two'], env={'BARRIER': str(barrier)})

    def test_e15_empty_success_and_usage(self):
        worker = self.worker('import sys\nprint(sys.argv[1])')
        self.ok(15, ['2', worker, '--'])
        self.ok(15, ['1', worker, '--', 'x'], b'x\n')
        self.bad(15, ['0', worker, '--', 'x'])
        self.bad(15, ['2', self.dir / 'missing', '--'])

    def test_e15_term_reaps_workers(self):
        ready = self.dir / 'ready'; ready.mkdir()
        worker = self.worker('import os, pathlib, signal, sys\np=pathlib.Path(os.environ["READY"])/sys.argv[1]\np.write_text(str(os.getpid()))\nsignal.pause()')
        process = self.launch(15, ['2', worker, '--', 'a', 'b'], {'READY': str(ready)})
        for name in ('a', 'b'): self.await_file(ready / name, process)
        pids = [int((ready / name).read_text()) for name in ('a', 'b')]
        process.send_signal(signal.SIGTERM)
        rc, _, _ = self.finish(process)
        self.assertEqual(rc, 143)
        for pid in pids:
            with self.assertRaises(ProcessLookupError): os.kill(pid, 0)

    def test_e15_relative_executable(self):
        worker = self.worker('import sys\nprint(sys.argv[1])', '-relative worker')
        self.ok(15, ['1', worker.name, '--', 'hello shell'], b'hello shell\n')

    def test_e16_literal_unique_sorted_and_no_match(self):
        path = self.file('log', b'z ERROR\na ERROR\nz ERROR\nX\n.*\n-n\n')
        self.ok(16, ['ERROR', path], b'a ERROR\nz ERROR\n')
        self.ok(16, ['.*', path], b'.*\n')
        self.ok(16, ['-n', path], b'-n\n')
        self.ok(16, ['missing', path])

    def test_e16_errors(self):
        self.bad(16, ['x', self.dir / 'missing'], 1)
        self.bad(16, ['x'])
        fake = self.dir / 'fake'; fake.mkdir()
        env = {'PATH': str(fake) + os.pathsep + self.env['PATH']}
        self.worker('import sys\nsys.exit(2)', 'fake/grep')
        self.bad(16, ['x', self.file('log', b'x\n')], 1, env=env)
        (fake / 'grep').unlink()
        self.worker('import sys\nsys.exit(9)', 'fake/sort')
        self.bad(16, ['x', self.dir / 'log'], 1, env=env)


if __name__ == '__main__':
    exercise = OPTIONS.exercise.lower()
    if exercise != 'all' and exercise not in {'e%02d' % i for i in range(1, 17)}:
        PARSER.error('exercise must be E01..E16 or all')
    names = unittest.defaultTestLoader.getTestCaseNames(WorkshopTests)
    selected = [name for name in names if exercise == 'all' or name.startswith('test_' + exercise + '_')]
    print('mode=%s | bash=%s | exercise=%s' % (OPTIONS.mode, OPTIONS.bash, OPTIONS.exercise), flush=True)
    suite = unittest.TestSuite(WorkshopTests(name) for name in selected)
    result = unittest.TextTestRunner(verbosity=2).run(suite)
    sys.exit(0 if result.wasSuccessful() else 1)
