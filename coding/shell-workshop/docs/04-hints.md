# 04 · 分层提示

先尝试 10～15 分钟，再展开第一层。只有能解释第一层仍写不出来时，才看第二层。完整实现放在 `solutions/`，本页不直接给出整段答案。

## E01 · 参数边界

<details><summary>提示 1：不要先拼接参数</summary>

参数本来就是一组有边界的值。用 `$#` 取得数量，使用 `for arg in "$@"` 遍历，空串也会成为一次迭代。

</details>
<details><summary>提示 2：格式与数据分开</summary>

先输出 `count=%s\n`，循环中用 `'%s=<%s>\n'` 作为固定格式串。计数从 1 开始，用 `i=$((i + 1))` 更新。不要把参数值写进格式串。

</details>

## E02 · 默认配置

<details><summary>提示 1：考虑空字符串</summary>

`${APP_ENV-dev}` 和 `${APP_ENV:-dev}` 对空串的表现不同。题目要求“未设置和空串都回退”。

</details>
<details><summary>提示 2：先校验入口</summary>

先判断 `$#` 是否为 0，错误写 stderr 并退出 2，然后用带引号的 `${APP_ENV:-dev}` 和 `${APP_PORT:-8080}` 作为 printf 参数。

</details>

## E03 · 整数校验

<details><summary>提示 1：先让算术输入变得可信</summary>

先检查参数数目，再限制为 1～3 位 ASCII 数字，之后才能做转换。限制长度还能避免超长值溢出。

</details>
<details><summary>提示 2：八进制不是你要的</summary>

使用 `[[ $1 =~ ^[0-9]{1,3}$ ]]`，再用 `$((10#$1))` 转换，检查不超过 255。求和可循环，也可用 `n*(n+1)/2`；本题上限足够小。

</details>

## E04 · 选项解析

<details><summary>提示 1：先收集，后输出</summary>

用 while/case，每读一个带值选项 shift 2。记录是否出现过 name/times，既判断必填，也拒绝重复。

</details>
<details><summary>提示 2：缺值检查要在读取 $2 前</summary>

`[ "$#" -ge 2 ]` 检查剩余参数数目。`--name -bot` 合法，不能以“下一个以 - 开头”判断缺值。遇到 `--` 后，剩余参数必须为零。一个统一的 usage 函数负责 stderr 和 exit 2。

</details>

## E05 · 完整读取最后一行

<details><summary>提示 1：read 的失败不意味着没有读到字符</summary>

EOF 前读到了内容但没有遇到换行时，read 会给变量赋值，却返回非零。需要在条件里检查变量是否仍有内容。

</details>
<details><summary>提示 2：三个保护组合</summary>

`while IFS= read -r line || [ -n "$line" ]; do ...; done < "$file"`。IFS 防止丢空白，-r 保留反斜杠，最后的条件补上无换行末行。打印之前计数加 1。

</details>

## E06 · 日志聚合

<details><summary>提示 1：决定什么时候允许输出</summary>

如果每读到好行就输出，一旦后面出现坏行，无法撤回 stdout。先累计在 awk 数组里，最后统一判断是否有 bad 标志。

</details>
<details><summary>提示 2：两个阶段都要检查</summary>

awk 用 `$3` 作为 service 键、`$4 >= 500` 判断错误计数。END 中先判断 bad，再输出 TAB 分隔结果。交给 C locale 的 sort 后，立即复制 PIPESTATUS，任一阶段失败就返回 1。

</details>

## E07 · 文件清单

<details><summary>提示 1：让 find 直接产生正确格式</summary>

不需要先取列表再逐个拼接：`find ROOT -type f -name '*.log' -print0` 已有需要的筛选与 NUL 分隔。默认不跟随符号链接。

</details>
<details><summary>提示 2：相对目录可能以 - 开头</summary>

把相对路径加上 `./`，避免被 find 当表达式/选项。根本身是符号链接要明确拒绝，检查时还要考虑路径结尾的 `/`。不要用 `result=$(find ...)` 保存 NUL。

</details>

## E08 · 命令执行

<details><summary>提示 1：参数已经是一份执行计划</summary>

移除 mode 和 `--` 后，剩余的 `"$@"` 就是命令及参数。dry-run 对它逐项打印，run 直接执行它。

</details>
<details><summary>提示 2：可以替换当前进程</summary>

`exec "$@"` 自动把命令的退出状态交给调用方。`printf '%q\n' "$arg"` 展示参数，不能把这个展示结果再送进 eval。

</details>

## E09 · 配置解析

<details><summary>提示 1：配置不是代码</summary>

用 read 读取，按第一个等号分开。`${line%%=*}` 得到等号之前的键，`${line#*=}` 去掉第一个等号及其前缀，保留后面其他等号。

</details>
<details><summary>提示 2：把校验和发布结果分开</summary>

记录两个键的 seen 标志；遇到重复或未知键就失败。PORT 限制长度后做十进制转换和范围校验。读完检查两个键都有，最后才 printf。`$(...)` 留在 value 中就是普通文本，除非你错误地用了 source/eval。

</details>

## E10 · 原子写入

<details><summary>提示 1：不要先打开目标写入</summary>

先验证目标类型，在目标父目录 mktemp，写入临时文件成功后才 mv 到目标。创建临时文件之前设 umask 077。

</details>
<details><summary>提示 2：关键状态不能靠直觉</summary>

`cat > "$tmp" || exit 1`，然后 `mv -f -- "$tmp" "$target" || exit 1`。临时文件创建成功后装 EXIT trap。先给相对目标加 `./`，可以减少前导 - 和没有斜杠时拆 parent 的问题。

</details>

## E11 · 备份

<details><summary>提示 1：备份先是一个未发布的产物</summary>

不要直接写最终文件名。先规范化源目录与归档父目录的物理路径，拒绝归档在源中，再生成临时归档并校验。

</details>
<details><summary>提示 2：检查“不存在”后仍有竞态</summary>

`tar -czf "$tmp" -C "$source" .` 包括隐藏文件。`tar -tzf` 检查可列出。要在目标已存在时失败，可以用同目录硬链接 `ln -- "$tmp" "$archive"` 发布，再删临时名字。本题假设父目录可信且目标不会被竞争式替换为目录。

</details>

## E12 · 重试

<details><summary>提示 1：成功与失败退出点分开</summary>

每轮先执行；成功立即退出；失败在 else 第一行存 `$?`。最后一次失败直接返回保存的状态，不再 sleep。

</details>
<details><summary>提示 2：只传递原始参数</summary>

移除次数、间隔和 `--` 后，在循环中运行 `"$@"`。不要用 `if ! "$@"` 后的 `$?` 当原状态；那里得到的是取反后的结果。

</details>

## E13 · 目录生命周期

<details><summary>提示 1：先拥有资源，再注册清理</summary>

mktemp 成功后才能保证临时目录属于你。EXIT 负责清理；TERM/INT 处理器负责让任务结束并选择退出状态。

</details>
<details><summary>提示 2：保存 PID，转发，再 wait</summary>

用 `WORK_DIR=$tmp "$@" <&0 &` 启动，保存 `$!`。主流程 wait 并保存状态；信号处理器向这个直接子进程发 TERM，wait 后 exit 143。处理信号时可暂时忽略后续 INT/TERM，避免清理重入。不要承诺管理整个后代进程树。

</details>

## E14 · 锁所有权

<details><summary>提示 1：把 mkdir 本身当判定</summary>

不要先判断目录不存在。mkdir 成功才拥有锁；失败返回 75，不能装会删除已有锁的 EXIT trap。

</details>
<details><summary>提示 2：临界区覆盖整个任务</summary>

成功后 trap 中用 rmdir 释放，只删除预期的空锁。命令退出后先保存状态，再退出。TERM 时先停止并 wait 子任务，然后才能解除互斥，否则下一任务可能与仍运行的前任务重叠。

</details>

## E15 · 并发上限

<details><summary>提示 1：先写容易证明的批处理</summary>

每启动一个后台任务记录 PID；数组长度达到 LIMIT 就逐个 wait 清空；循环结束后再回收尾批。这满足上限，兼容 Bash 3.2。

</details>
<details><summary>提示 2：失败不应该阻止回收</summary>

用 failed 标记累计失败；每个 wait 都执行，不要第一个失败就 exit。信号处理器遍历当前 PID 列表发送 TERM，再逐个 wait。worker 是文件路径，相对路径要明确带 `./`，不能误交给 PATH 查找。子任务 stdin 使用 `/dev/null`。

</details>

## E16 · 管道状态

<details><summary>提示 1：非零不只有一种含义</summary>

grep 0 表示匹配，1 表示无匹配，其他状态才是本题的操作失败。sort 则需要 0。不能只看最后一段，也不能无条件忽略管道错误。

</details>
<details><summary>提示 2：复制 PIPESTATUS 必须及时</summary>

管道之后紧接 `statuses=("${PIPESTATUS[@]}")`；之后再输出日志或判断。固定字符串用 `grep -F -e "$1"`，文件可用重定向输入，避免路径被当成选项。

</details>

## 排错题的方向提示

| 题目 | 先找的根因 |
|---|---|
| D01 | 一份文件名列表被经过了多少次不合适的分割？ |
| D02 | 字符串里还保存着原来的两个参数边界吗？ |
| D03 | `$?` 指向原命令，还是 `!` 之后的条件？ |
| D04 | Shell 何时执行 `>`？ |
| D05 | grep 1、grep 2、sort 非零分别是什么含义？ |
| D06 | `mkdir -p` 成功是否意味着目录是这次新建的？ |
| D07 | 为什么关键词中的分号会从数据变成语法？ |
| D08 | 清理逻辑是否知道当前进程真的拥有锁？ |
