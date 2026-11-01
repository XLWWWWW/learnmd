# 01 · Shell 编程详解

本章默认使用 Bash。标为“片段”的代码用于理解某个机制；涉及文件的片段需要先准备对应文件。可直接运行的完整演示在 `examples/`，完整工作脚本在 `solutions/`。不要把含 `exit` 的教学片段直接粘进正在工作的交互终端；保存成脚本再用 `bash 文件名` 运行。

## 1. Shell 到底负责什么

终端是交互窗口，Shell 是解释命令的程序，`ls`、`tar`、`curl` 等是它调用的工具。一个 Shell 脚本通常不亲自实现复杂算法，而是负责把程序组合起来：准备参数、连接输入输出、检查退出码、安排顺序、处理失败。

例如：

```bash
grep -F 'ERROR' app.log | sort | uniq -c
```

Shell 建立管道；grep 选行，sort 排序，uniq 统计相邻重复行。数据主要在程序间流动，不必全部放进变量。脚本适合部署步骤编排、文件整理、日志摘要、定时维护、开发工具包装。复杂 JSON 变换、大型状态机、严格类型校验、长期运行服务，通常更适合 Python/Go 等；Shell 可以做它们的薄入口。

## 2. Bash、sh、zsh 与脚本启动

`sh` 表示一种标准化 Shell 接口，但不同系统的 `/bin/sh` 可能是不同实现。Bash 支持更多功能；zsh 又有自己的展开和数组语义。不能因为在 zsh 终端里试过一行命令，就认为 Bash 脚本表现相同。

```bash
#!/usr/bin/env bash
printf 'hello shell\n'
```

保存为 `hello.sh`，使用 `chmod +x hello.sh`，然后 `./hello.sh`。操作系统依据第一行 shebang 选解释器。`bash hello.sh` 明确交给 Bash；`sh hello.sh` 明确交给 sh，不能靠 shebang 把它变回 Bash。

| 写法 | 意义 |
|---|---|
| `./script.sh` | 新进程；依据 shebang；需要执行权限 |
| `bash script.sh` | 新 Bash；需要文件可读；不需要脚本执行位 |
| `source script.sh` 或 `. script.sh` | 在当前 Shell 执行；变量、cd、set、exit 都会影响当前环境 |
| `exec command args...` | 用 command 替换当前 Shell 进程；不会返回继续下一行 |

`#!/usr/bin/env bash` 用 PATH 查找 Bash，便利但版本取决于环境；`#!/bin/bash` 固定路径，macOS 上通常是较旧版本。本工程实际测试使用 `/bin/bash`，可通过 `BASH_BIN` 改选。

## 3. 一条命令的结构与退出状态

```bash
APP_ENV=dev command arg1 'arg two' >out.txt 2>err.txt
```

命令前的赋值进入这次执行的环境；两个参数分别是 `arg1` 和 `arg two`；重定向由 Shell 设置，文件名不是 command 的普通参数。多数外部命令可以接收 `--` 结束选项，但要查看该命令自己的协议。

每条命令返回 0～255 的退出状态。**0 表示成功，非 0 需按命令语义解释。** grep 的 1 是“没有匹配”，不等于文件读取失败。信号退出常由 Shell 表示成 `128 + 信号编号`，不能把所有非零都粗暴归类为同一种错误。

```bash
if command_that_may_fail; then
    printf 'success\n'
else
    rc=$?                       # 第一时间保存；后续命令会改变 $?
    printf 'failed: %s\n' "$rc" >&2
fi
```

`&&` 只在左边成功时执行右边，`||` 只在左边失败时执行右边。`a && b || c` 不是通用三元表达式：a 成功但 b 失败时也会执行 c。需要明确的两个分支就写 `if`。

## 4. 变量、环境与参数

```bash
name='hello shell'               # 等号两边不能加空格
printf '%s\n' "$name"
export APP_ENV=dev               # 后代进程可以读取
APP_ENV=test bash -c 'printf "%s\n" "$APP_ENV"'
```

Shell 变量通常是字符串；export 赋予变量“传给后代进程”的属性。子进程改自己的环境，不会反向修改父 Shell。用 `source` 读取配置会执行其中所有代码，所以不适合不可信或只应作为数据的配置。

| 参数 | 含义 |
|---|---|
| `$0` | 脚本调用名；不保证是绝对路径 |
| `$1`、`${10}` | 第 1、第 10 个参数；`$10` 是 `$1` 后跟字符 0 |
| `$#` | 参数数量 |
| `"$@"` | 保留每个参数边界，零个参数就展开成零个参数 |
| `"$*"` | 用 IFS 的第一个字符连接所有参数，形成一个参数 |
| `$?` | 最近完成的命令或管道的状态 |
| `$$`、`$!` | Shell 的 PID；最近一次后台任务的 PID |

函数也有自己的位置参数；`shift` 丢弃最前面的参数，通常用于解析命令行。不要用 `HOME`、`PATH`、`IFS`、`SHELL` 等有系统含义的变量名存放普通业务数据。

## 5. 引号与展开：最值得花时间理解的部分

Shell 先解析语法，然后进行多种展开。对日常脚本最有影响的顺序是：变量/命令/算术展开 → 未引用结果的单词分割 → 文件名匹配。并不是把变量内容当作一段新的 Shell 源码重新解释。

```bash
value='a b *'
printf '<%s>\n' "$value"         # 一个参数：a b *
printf '<%s>\n' $value           # 先分成 a、b、*，然后 * 还会匹配当前目录
printf '<%s>\n' '$value'         # 字面量 $value
```

- 单引号保留字面内容，其中不能直接放单引号。
- 双引号保留变量与命令替换，但阻止普通的分割和通配符展开。
- `\` 可以转义单个字符；Bash 的 `$'a\nb'` 会产生真正的换行。
- `"$dir"/*.log` 引用目录部分，允许最后的 glob；`"$dir/*.log"` 把星号也变成字面量。
- 给变量使用引号，不能阻止目标命令把 `-rf` 解释成选项；这由 `--` 或明确的路径前缀解决。

`cmd='printf hello'` 再运行 `"$cmd"`，会寻找一个名字叫 `printf hello` 的程序；运行 `$cmd` 又会依赖分割。正确方式是数组或函数，见第 14 节。

关于展开顺序见 [Bash Shell Expansions](https://www.gnu.org/s/bash/manual/html_node/Shell-Expansions.html)；关于未引用参数的典型问题见 [ShellCheck SC2086](https://www.shellcheck.net/wiki/SC2086)。

## 6. 默认值、字符串处理与数字

| 写法 | 作用 |
|---|---|
| `${x:-fallback}` | x 未设置或为空时使用 fallback，不修改 x |
| `${x-fallback}` | 只有未设置时使用 fallback，空串仍保留 |
| `${x:=fallback}` | 未设置或为空时给 x 赋默认值 |
| `${x:?message}` | 未设置或为空时报错；非交互 Shell 退出 |
| `${x:+yes}` | x 已设置且非空时展开成 yes，否则为空 |
| `${#x}` | 字符串长度；字符/字节关系受 locale 影响 |
| `${x#pattern}` / `${x##pattern}` | 去掉最短/最长的匹配前缀 |
| `${x%pattern}` / `${x%%pattern}` | 去掉最短/最长的匹配后缀 |

```bash
file='/var/log/app.error.log'
printf '%s\n' "${file##*/}"      # app.error.log
printf '%s\n' "${file%.*}"       # /var/log/app.error
```

这里的 pattern 是 glob 模式，不是正则表达式。手工拆路径还要考虑根目录、尾斜杠、没有斜杠等特殊情况；不要当成完整路径规范化算法。

整数计算用 `$((...))`，不会计算浮点数：`$((5 / 2))` 是 2。以 0 开头的数可能按八进制解释。对外部输入先校验字符、长度和范围，再用 `10#` 指定十进制；不要直接把任意文本送进算术表达式。

```bash
raw=008
if [[ $raw =~ ^[0-9]{1,3}$ ]]; then
    n=$((10#$raw))
    printf '%s\n' "$((n + 1))"   # 9
fi
```

`((expression))` 本身是命令：表达式值非 0 时状态为 0，值为 0 时状态为 1。因此 `i=0; ((i++))` 的退出状态是 1。计数时写 `i=$((i + 1))` 可以避免与 `set -e` 发生意外互动。

进一步阅读：[Bash 参数展开](https://www.gnu.org/s/bash/manual/html_node/Shell-Parameter-Expansion.html)。

## 7. 条件判断：test、[[ ]]、glob 与正则

```bash
if [ -f "$file" ] && [ -r "$file" ]; then
    printf 'readable regular file\n'
fi
if [[ $name == api-* ]]; then printf 'api service\n'; fi
if [[ $port =~ ^[0-9]+$ ]]; then printf 'digits\n'; fi
```

`[` 是命令，最后的 `]` 是参数，空格必须写。`[[ ... ]]` 是 Bash 语法，避免了很多普通参数分割问题，支持模式匹配和正则，但不是通用 POSIX sh 语法。

| 判断 | 常见含义 |
|---|---|
| `[ -z "$s" ]` / `[ -n "$s" ]` | 空 / 非空字符串 |
| `[ "$a" = "$b" ]` | 字符串相等 |
| `[ "$a" -eq "$b" ]` | 整数相等，需要先保证输入是整数 |
| `[ -e "$p" ]` / `[ -f "$p" ]` / `[ -d "$p" ]` | 存在 / 普通文件 / 目录 |
| `[ -L "$p" ]` | 路径本身是符号链接；断链也能被发现 |
| `[ -r "$p" ]` / `[ -x "$p" ]` | 当前身份可读 / 可执行 |
| `[[ $s == pre* ]]` | glob 前缀匹配；`"pre*"` 则是字面量 |
| `[[ $s =~ ^[0-9]+$ ]]` | 扩展正则；把整个右侧引用会改变匹配含义 |

glob 的 `*` 表示任意字符串，正则的 `*` 表示前一个单元重复。正则匹配任意字符串通常是 `.*`。`-f` 检查普通文件时会跟随符号链接；需要拒绝链接就额外检查 `! -L`。

检查结果和实际使用之间可能发生变化，称为检查与使用之间的竞态。`[ ! -e target ]` 不是一个并发创建锁，也不能独立保证后续不会覆盖文件。

## 8. 分支、循环与函数

```bash
case ${1:-help} in
    start|stop) printf 'action=%s\n' "$1" ;;
    help) printf 'usage: tool start|stop\n' ;;
    *) printf 'unknown action\n' >&2; exit 2 ;;
esac

for arg in "$@"; do
    printf '%s\n' "$arg"
done

log() {
    local level=$1
    shift
    printf '[%s] %s\n' "$level" "$*" >&2
}
```

`case` 适合命令分发与有限枚举。`for ((i=0; i<3; i++))` 是 Bash 算术循环；`while command; do ...; done` 在 command 成功时继续；`until` 在失败时继续。`break` 退出循环，`continue` 跳到下一轮。

函数用 stdout 返回数据，用 `return` 返回状态，不用 `return 12345` 返回业务数字。Bash 的 local 是动态作用域，子函数也可能看到调用函数的局部变量，函数间仍应明确传参。写 `local value; value=$(command) || return`，避免 `local value=$(command)` 用 local 的成功状态掩盖 command 的失败。

## 9. stdin、stdout、stderr 与重定向

进程通常有三个标准描述符：0 是输入，1 是结果输出，2 是诊断输出。把结果与日志分开，调用者才可以安全地保存结果。

| 写法 | 含义 |
|---|---|
| `< file` | 从文件输入 |
| `> file` / `>> file` | 截断写 / 追加写 stdout |
| `2> file` | 将 stderr 写入文件 |
| `>/dev/null` | 丢弃 stdout，不丢弃 stderr |
| `>file 2>&1` | stdout 去 file，然后 stderr 复制当前 stdout 的去向 |
| `2>&1 >file` | stderr 先复制原 stdout，随后只把 stdout 改去 file |

重定向按从左到右顺序建立。`command > important.txt` 会在 command 真正执行前截断文件，即使命令不存在，也可能已经把原内容清空。需要完整成功后更新文件，使用同目录临时文件再发布，见 E10。

Here-document 适合多行文本：

```bash
cat <<'EOF'
$HOME 不展开，$(date) 也不执行。
EOF
```

去掉分隔符的引号会允许正文中的参数展开、命令替换和算术展开。`<<-EOF` 只剥离开头的 TAB，不剥离普通空格。Bash 的 `<<< "$text"` 是 here-string，会追加换行，不是原样二进制输入。

## 10. 管道、子 Shell 与命令替换

管道连接 stdout 和下一程序的 stdin；各阶段通常并发运行。stderr 默认仍去原来的 stderr。Bash 默认取最后一个阶段的状态；`set -o pipefail` 改为最右侧非零阶段的状态，而不是按时间上最后失败的阶段。

```bash
false | true
printf '%s\n' "$?"              # 默认 0
set -o pipefail
false | true
printf '%s\n' "$?"              # 1
```

如需区分每个阶段，要立即复制数组：

```bash
grep -F 'ERROR' app.log | sort -u
statuses=("${PIPESTATUS[@]}")    # 中间不要插入 echo、赋值或其他命令
```

默认 Bash 管道阶段在子 Shell 环境中运行；不要依赖最后一段修改外层变量：

```bash
count=0
printf 'a\nb\n' | while IFS= read -r line; do count=$((count + 1)); done
printf '%s\n' "$count"          # 本工程默认环境下仍是 0
```

更直接的文件读取方式是 `while ...; do ...; done < file`。`done < <(producer)` 通过进程替换把 producer 的输出当输入，循环在当前 Shell 运行；但 producer 的失败不会自动成为循环退出码。关键生产步骤最好先检查成功再消费，或者明确管理进程状态。较新 Bash 的 `lastpipe` 还会影响管道末段行为，本教程不依赖它。

`( commands )` 创建子 Shell 环境；`{ commands; }` 在当前环境分组。`cd`、普通变量修改等不会从子环境传回。

```bash
text=$(printf 'a\n\n')           # 命令替换去掉所有尾部换行，结果是 a
```

即使写 `"$(...)"` 也保不住这些尾部换行。Bash 变量不能保存 NUL；二进制和 NUL 清单要通过文件或流传递。

## 11. 正确读行与打印

```bash
while IFS= read -r line || [ -n "$line" ]; do
    printf '%s\n' "$line"
done < input.txt
```

`IFS=` 防止 read 去掉首尾分隔空白，`-r` 防止反斜杠被当作转义，`|| [ -n "$line" ]` 处理最后一行没有换行的情况。空文件没有循环；一个只含换行的文件有一行空内容。这里处理的是文本行，不是任意二进制。

`for x in $(cat file)` 会同时丢失空行、空格和参数边界。`echo` 对 `-n`、反斜杠等的行为受实现和选项影响，稳定输出优先用 `printf`。

**格式串应当由代码固定**：`printf '%s\n' "$input"`。不要写 `printf "$input"`，否则输入中的 `%s`、`%n`、反斜杠等可能被解释为格式控制。

## 12. 文本工具怎么分工

| 工具 | 适合做什么 | 常见陷阱 |
|---|---|---|
| `grep -F` | 字面字符串筛选 | 没有匹配返回 1；不代表操作错误 |
| `grep -E` | 扩展正则筛选 | 正则与 glob 不同 |
| `sed` | 简单替换和行变换 | `sed -i` 的 macOS/Linux 形式不同 |
| `awk` | 按字段计算、分组计数 | 默认空白分字段；不是 CSV/JSON 通用解析器 |
| `sort` | 按规则排序 | 默认文本排序；数字通常需要 `-n` |
| `uniq -c` | 统计连续重复行 | 不先排序只能统计相邻重复 |
| `cut` | 固定分隔符选字段 | 不理解 CSV 引号内的逗号 |
| `tr` | 字符映射/删除 | 不处理通用正则或结构化记录 |
| `wc -l` | 统计换行数 | 没有末尾换行的文本行不会额外计数 |

```bash
awk '$4 >= 500 { count[$3]++ } END { for (s in count) print s, count[s] }' access.log
```

这是简化格式的片段，不适用于任意 Nginx 日志；E06 明确定义输入并校验格式。awk 的关联数组迭代顺序不保证，要求稳定输出时排序。`LC_ALL=C sort` 使排序可预测，但这不是自然语言排序。

处理 JSON 使用 jq 或真正的 JSON 库；处理带引号、嵌入逗号/换行的 CSV 使用 CSV 库。解析格式应由格式规范决定，不应因为手边只有正则就勉强拼出来。

## 13. 文件名与 NUL 清单

Unix 文件名可以有空格、TAB、换行、星号和前导 `-`，但路径组件不能包含 `/` 或 NUL。所以“每行一个文件名”不是通用协议。

```bash
find ./logs -type f -name '*.log' -print0 |
while IFS= read -r -d '' file; do
    printf 'found: %q\n' "$file"
done
```

`-print0` 与 Bash 的 `read -d ''` 配合，用 NUL 分隔。这里循环只打印，不要求在外层保留变量；如果需要保留，使用已成功生成的清单文件再重定向读取。find 的 `-name '*.log'` 要引用，让模式交给 find，而不是先被当前 Shell 展开。

尽量不解析 `ls` 输出。执行同一个命令处理文件时，可考虑 `find ... -exec command {} +`，参数边界由 find 维护；注意命令的选项协议。`xargs -0` 可以消费 NUL，但 GNU 与 BSD 的空输入、并发和扩展选项要单独检查。

默认 glob 不匹配点开头文件；无匹配时 Bash 通常保留字面模式。可以明确设置 `shopt -s nullglob`，或在循环中检查 `[ -e "$file" ] || [ -L "$file" ]`。修改 shopt 会影响之后的代码，库函数应恢复原状态或放在子 Shell 中。

## 14. 数组、选项解析与命令封装

```bash
args=(--connect-timeout 3 --max-time 10)
url='https://example.com/path?a=1&b=2'
args+=(--url "$url")
# 完整执行片段：需要安装 curl，并替换成你拥有的服务地址。
# curl "${args[@]}"
```

Bash 索引数组在本工程支持范围内。`"${args[@]}"` 保留每个元素，`"${args[*]}"` 连接成一个字符串。Bash 关联数组 `declare -A` 需要 Bash 4+，不能在 macOS 自带 3.2 上直接使用。`mapfile/readarray`、`wait -n` 等也不属于本工程的基础版本。

短选项适合用内建 `getopts`：

```bash
while getopts ':n:v' opt; do
    case $opt in
        n) name=$OPTARG ;;
        v) verbose=1 ;;
        :) printf 'option needs a value\n' >&2; exit 2 ;;
        \?) printf 'unknown option\n' >&2; exit 2 ;;
    esac
done
shift "$((OPTIND - 1))"
```

长选项可用 `while/case/shift` 明确解析（E04）；外部 `getopt` 的 GNU/BSD 行为有差异，不要默认同一种实现。名字以 `-` 开头的“选项值”如何解释、是否允许重复选项，都应写在接口契约里。

不要用 `eval "$command"` 解决参数拼接问题。eval 会把数据重新作为代码解析，命令中的分号、`$(...)` 等随即有了执行意义。使用 `"$@"`、数组和函数传递参数。

## 15. 错误处理：理解 set，再决定使用方式

常见的 `set -euo pipefail` 分别开启 errexit、未设置变量报错、管道失败传播；它们不会自动证明脚本可靠。

- `-e` 有语法上下文例外：if/while 的条件、部分 `&&/||` 列表、`!`、管道阶段等都可能抑制退出。函数被当作条件调用时，函数内部的失败行为也会受影响。
- `-u` 要配合 `${1:-}`、`${OPTIONAL:-}` 处理可选输入；它不是数据校验。旧 Bash 对空数组等还有版本细节。
- `pipefail` 把 grep 的“没有匹配”也视为管道非零；`producer | head` 还可能因消费者提前关闭而产生 SIGPIPE。要结合业务语义判断。
- `ERR` trap 同样有条件例外，不等同于 Python/Java 的全局异常捕获。

```bash
if download_step; then
    printf 'download complete\n' >&2
else
    rc=$?
    printf 'download failed: %s\n' "$rc" >&2
    exit "$rc"
fi
```

本工程参考答案选择显式检查关键状态，没有机械地给所有脚本加 `set -e`。这不意味着不能使用它，而是让失败协议在练习中清晰可见。开启 strict 选项后仍应测试“预期非零”和“真实故障”。

准确规则见 [Bash set 内建命令](https://www.gnu.org/s/bash/manual/html_node/The-Set-Builtin.html)。

## 16. 临时文件、原子发布与备份

`/tmp/result.$$` 的名字可预测，可能冲突或被抢先创建。使用 `mktemp` 创建独有文件/目录；用 `umask 077` 限制默认权限；创建成功后再安装清理逻辑。

原子替换的流程是：在目标同目录建立临时文件 → 写入并检查 → `mv` 发布 → 清理残留。读者通常看到旧内容或新内容，不会看到一半写入的目标。这只解决“发布可见性”，不等于磁盘断电后的持久性，也不能自动解决多个写者互相覆盖。

E10 约定可信的本地父目录、单写者，拒绝符号链接和目录目标；新文件权限是 0600，不保留旧文件的权限、属主、ACL 或扩展属性。生产配置替换需要根据实际服务设置元数据、验证内容并协调 reload。

备份也应先写临时归档，确认 tar 成功后再发布。归档不能放在正在归档的源目录里。**能列出归档不等于恢复验证**；至少把自己生成的可信备份解压到独立目录并比对。正在持续变化的数据源可能需要数据库专用备份或文件系统快照。

## 17. trap、信号与清理

```bash
tmp=$(mktemp -d "${TMPDIR:-/tmp}/mytask.XXXXXX") || exit 1
trap 'rm -rf -- "$tmp"' EXIT
trap 'exit 130' INT
trap 'exit 143' TERM
```

这个片段只删除自己创建的临时目录。EXIT 在正常退出、显式 exit 等场景触发；INT 常来自 Ctrl-C，TERM 常来自停止任务。SIGKILL（kill -9）、机器断电等不能被 trap 捕获，因此清理不能只有一个理想路径。

安装 trap 时的引号决定变量何时展开：单引号通常让表达式在 trap 执行时求值。清理函数不能把原来的业务失败覆盖成成功，也不能删除没有由自己成功获得的锁。

Bash 等待前台外部程序时，trap 执行可能延后；使用后台进程加 `wait` 更方便主动响应，再把信号传给子进程并等待回收。E13/E14 展示的是**直接子进程**管理；它要求子命令不 daemonize、不再派生需要单独管理的后代，并响应 TERM。可靠管理完整进程树、更强超时和强杀兜底应交给任务管理器或系统服务管理机制。

## 18. 后台任务、wait 与有界并发

```bash
task_a &
pid_a=$!
task_b &
pid_b=$!
failed=0
wait "$pid_a" || failed=1
wait "$pid_b" || failed=1
exit "$failed"
```

`&` 只是启动，不代表任务成功；要保存 PID 并逐一 wait。裸 `wait` 等全部任务，不能替你给出完整的逐项失败表。后台任务读取共享 stdin 会互相抢输入，所以批处理一般显式重定向为 `/dev/null`。

不要对十万项直接全部后台启动。E15 用“每批最多 N 个，等待本批，再开下一批”兼容 Bash 3.2。这是真正有上限的并发，但利用率不如动态补位的工作池。多任务 stdout/stderr 可能交错；需要稳定报表时每项写独立结果文件，再串行合并。

## 19. 互斥、幂等与重试

用 `if [ ! -d lock ]; then mkdir lock; fi` 抢锁存在竞态。直接 `mkdir lock`，成功者才拥有锁；失败者不执行任务、不删除别人的锁。E14 用这个模式教学。Linux 上的 `flock` 可以简化文件锁，但本机没有这个命令也能完成练习。

锁必须定义：锁覆盖哪段操作、失败如何退出、异常后是否有陈旧锁、由谁判断可清理。目录锁在 SIGKILL 后会遗留，不能仅凭“很旧”就自动删除；PID 也可能被复用。对网络文件系统和跨机器协调要另做方案。

幂等意味着重复执行不会重复产生不可接受的副作用。覆盖同样的缓存通常更容易幂等；转账、发消息、创建订单则需要幂等键或事务支持。重试前先问操作是否能安全重复。

E12 使用有限次数和固定间隔，全部失败时保留最后一次状态。生产网络调用还应分类可重试错误、设置单次超时与总预算，必要时指数退避加随机抖动；不要把权限不足、参数错误也无限重试。

## 20. 调试、测试与跨平台

```bash
bash -n script.sh                # 只检查语法
bash -x script.sh arg            # 打印展开后的执行轨迹；可能暴露口令
./practice.sh check
./practice.sh solution
```

诊断顺序：确定解释器 → 确定 cwd/PATH/locale → 查看实际参数 → 保存退出码 → 检查 stderr → 缩小成最小复现。`set -x` 不应在含 token/password 的任务上长期启用。

| 项目 | 需要注意的差异 |
|---|---|
| `sed -i` | GNU 与 macOS BSD 对备份后缀参数的写法不同；临时文件发布更易统一 |
| `date` | GNU `date -d` 与 BSD `date -v` 不通用 |
| `stat` | GNU `-c` 与 BSD `-f` 的格式选项不同 |
| `readlink -f` | 不能假设所有 macOS/Unix 都支持；本工程不依赖 |
| `timeout` / `flock` | 常见于 Linux，macOS 基础环境未必有 |
| `find`、`xargs` | `-printf`、`-r` 等扩展的可用性不同 |
| Bash 新特性 | 关联数组、mapfile、wait -n、大小写转换展开要检查版本 |
| cron/CI | 工作目录、PATH、环境变量和交互终端不同 |

本项目的黑盒测试覆盖正常输入、特殊字符、失败注入、状态传播、锁争用和信号清理。它们证明这些约定用例在当前环境中成立，不证明任意文件系统和并发条件都安全。将脚本带到 Linux/容器时再运行同一套测试。
