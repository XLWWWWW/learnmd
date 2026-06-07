# Anthropic HH-RLHF 论文解读：Helpful and Harmless Assistant

## 1. 论文背景

Anthropic 的论文 **Training a Helpful and Harmless Assistant with Reinforcement Learning from Human Feedback** 发布于 2022 年，是早期系统研究 RLHF 对齐语言模型的重要论文之一。论文目标是训练一个既有帮助又尽量无害的对话助手。

这篇论文和 OpenAI 的 InstructGPT 非常接近，都属于基于人类反馈的强化学习路线，也就是 RLHF。两者共同点是：先让模型生成回答，再让人类比较回答优劣，训练一个偏好模型或奖励模型，最后用强化学习优化生成模型。

但 Anthropic 这篇论文的重点更偏向：

- 如何把 assistant 训练成 helpful and harmless。
- 如何收集 helpfulness 和 harmlessness 两类偏好数据。
- 如何使用 red-teaming 数据让模型更安全。
- 如何通过 rejection sampling 和 online RLHF 迭代收集更高质量数据。
- 如何研究模型规模、RLHF 训练和安全性之间的关系。

## 2. 核心问题

基础语言模型的预训练目标通常是预测下一个 token。这个目标可以让模型学习语言、知识和生成能力，但它并不等价于“符合用户意图”。

在真实助手场景中，用户希望模型满足两个基本要求：

- Helpful：回答有帮助、具体、相关、清楚。
- Harmless：不鼓励违法、危险、欺骗、暴力、滥用或其他有害行为。

这两个目标有时会冲突。例如用户提出危险请求时，如果模型完全追求 helpful，可能会给出危险细节；如果模型过度追求 harmless，又可能对正常问题过度拒答。因此，论文要解决的问题不是简单让模型“更强”，而是让模型在 helpfulness 和 harmlessness 之间学会更合适的行为边界。

## 3. 整体训练流程

Anthropic 的训练流程可以概括为：

```text
预训练语言模型
    ↓
Context Distillation
    ↓
收集人类偏好比较数据
    ↓
训练 Preference Model
    ↓
用 Preference Model 作为奖励函数
    ↓
PPO / RLHF 优化助手模型
    ↓
Helpful and Harmless Assistant
```

其中最关键的是偏好比较数据。论文不是主要让标注员写标准答案，而是让标注员比较两个模型回答哪个更好。

基本数据单元是：

```text
对话上下文 / 用户问题
    ↓
模型回答 A
模型回答 B
    ↓
人类标注：A 更好，或 B 更好
```

这类数据被用于训练 Preference Model。训练好的 Preference Model 再作为 RLHF 阶段的奖励模型。

## 4. 模型设计

### 4.1 基础模型

论文中的模型基于 Transformer decoder-only 语言模型，和 GPT 系列类似，都是自回归语言模型。基础模型先通过大规模文本进行预训练，具备通用语言生成能力。

论文研究了不同参数规模的模型，覆盖从较小模型到更大模型的多个规模。作者关注的不是单一模型架构创新，而是模型规模、偏好模型和 RLHF 训练共同作用后的行为变化。

可以把模型分成三类：

- Base LM：基础预训练语言模型。
- Preference Model：用于评价回答好坏的偏好模型。
- RLHF Policy：经过 PPO 强化学习优化后的助手模型。

### 4.2 Context Distillation

Anthropic 使用了一个重要的初始步骤：**context distillation**。

基础语言模型本身并不天然知道自己应该扮演“有帮助、诚实、无害的助手”。一种简单方法是在 prompt 里写一长段系统说明，例如：

```text
你是一个 helpful, honest, harmless 的 AI assistant。
你应该回答用户问题，但避免有害内容。
```

如果每次对话都显式加入这种上下文，模型会更像助手。但这种方法有两个问题：

- 每次都要携带长上下文，成本高。
- 模型行为依赖 prompt，不够稳定。

context distillation 的思想是：先用带有长上下文说明的 prompt 让模型产生助手式回答，再通过监督学习把这种行为蒸馏进模型参数中。

简化理解：

```text
带 HHH 说明的模型行为
    ↓
生成训练样本
    ↓
监督微调
    ↓
不需要长说明也能表现得像 HHH assistant
```

这里的 HHH 指 helpful, honest, harmless。

context distillation 和 InstructGPT 的 SFT 有相似之处，都是让模型先具备初步指令/助手行为。但二者的数据来源和具体目标不完全相同：InstructGPT 的 SFT 主要依赖人工示范回答；Anthropic 的 context distillation 更强调把预设上下文中的 assistant 行为蒸馏进模型。

### 4.3 Preference Model

Preference Model 是 Anthropic 论文中的核心组件。它输入一个对话上下文和一个候选回答，输出一个标量分数，表示这个回答符合人类偏好的程度。

形式上：

```text
输入：conversation x + response y
输出：score r(x, y)
```

如果人类认为回答 A 比回答 B 更好，那么训练目标就是：

```text
r(x, A) > r(x, B)
```

Preference Model 的作用类似 InstructGPT 中的 Reward Model。两者本质上都是把人类偏好排序转化成模型可优化的 reward 信号。

## 5. 数据集设计

论文中数据集主要是人类偏好比较数据，也就是 comparison data。这里的 comparison 不是一条问答，而是一次“两个回答哪个更好”的比较。

论文中提到三类重要数据集：

- Core base dataset。
- RS dataset。
- Online dataset。

### 5.1 Core Base Dataset

Core base dataset 是基础偏好数据集，使用 context-distilled LM 生成候选回答，然后让人类进行比较标注。

该数据集包括：

- 约 44k helpfulness comparisons。
- 约 42k red-teaming / harmlessness comparisons。

其中 helpfulness comparisons 表示正常助手场景下的偏好比较。标注员判断两个回答中哪个更有帮助、更相关、更清楚、更符合用户需求。

red-teaming / harmlessness comparisons 表示安全性相关比较。red-team 场景通常包含诱导模型输出危险、违法、攻击性或不安全内容的用户请求。标注员需要判断回答的 harmfulness，训练时则用于让模型偏好更无害、更安全的回答。

可以理解为：

```text
正常用户问题
    ↓
比较哪个回答更 helpful

危险或对抗性问题
    ↓
比较哪个回答更 harmless
```

论文中还说明，一段 conversation 通常会产生大约 4 个 comparisons。也就是说，一个对话样本可能不是只对应一次比较，而可能产生多个回答对比较。

### 5.2 RS Dataset：Rejection Sampling 数据集

RS 是 Rejection Sampling，拒绝采样。这里的 RS dataset 是一种更高质量的数据收集方式。

流程如下：

```text
先用 core base dataset 训练一个初版 Preference Model
    ↓
对同一个 prompt 让模型生成多个候选回答
    ↓
用初版 Preference Model 给候选回答打分
    ↓
拒绝低分回答，保留高分回答
    ↓
把筛选后的回答交给人类比较
    ↓
得到新的 comparison 数据
```

论文中的 RS dataset 包括：

- 约 52k helpfulness comparisons。
- 约 2k red-teaming comparisons。

为什么需要 RS？

如果只用基础模型随机生成回答，很多回答质量很差，人类比较很容易。例如一个回答明显胡说，另一个回答基本正确，这种比较对训练强偏好模型的价值有限。

当模型变强后，更有价值的数据是：

```text
一个还不错的回答
vs
一个更好的回答
```

RS 的作用就是用已有 Preference Model 先筛掉明显差的回答，让人类比较更接近、更高质量、更难区分的候选回答，从而训练出更细致的 Preference Model。

### 5.3 Online Dataset：迭代在线数据集

Online dataset 是论文中另一个重要数据来源。这里的 online 不是简单指线上用户日志，而是指数据来自一个迭代训练过程中的 RLHF 模型。

流程如下：

```text
第 1 周：训练 RLHF 模型 v1
    ↓
用 v1 生成回答
    ↓
人类比较哪个回答更 helpful
    ↓
加入 online dataset

第 2 周：更新 RLHF 模型到 v2
    ↓
用 v2 生成回答
    ↓
继续收集人类 helpfulness 比较
    ↓
加入 online dataset

持续约 5 周
```

论文中提到，online dataset 包含：

- 约 22k helpfulness comparisons。
- 不包含 red-teaming data。

为什么需要 online dataset？

因为模型经过 RLHF 后，输出分布会变化。基础模型可能犯很明显的错误，而 RLHF 模型的回答往往都比较像样。此时 Preference Model 需要学会区分更细微的质量差异，例如：

- 两个回答都相关，但哪个更准确？
- 两个回答都安全，但哪个更有帮助？
- 一个回答更礼貌但信息少，另一个回答更具体但措辞一般，哪个更好？

Online dataset 的作用是让偏好模型持续看到当前 RLHF 模型真实会生成的回答，从而减少训练数据和当前模型行为之间的分布偏移。

## 6. 标注任务怎么理解

Anthropic 的标注主要不是“写标准答案”，而是“比较两个回答”。

### 6.1 Helpfulness 标注

helpfulness 标注关注回答是否有帮助。标注员通常比较：

- 哪个回答更准确？
- 哪个回答更具体？
- 哪个回答更符合用户问题？
- 哪个回答更清楚？
- 哪个回答更少废话？
- 哪个回答更像一个有用助手？

数据形式：

```text
prompt
response A
response B
label: A is better / B is better
```

### 6.2 Harmlessness / Red-teaming 标注

harmlessness 标注关注回答是否安全。red-teaming prompts 通常会故意诱导模型输出有害内容。

标注员关注：

- 哪个回答更危险？
- 哪个回答更可能帮助用户实施伤害？
- 哪个回答更不应该被助手输出？
- 哪个回答拒绝得更合理？
- 哪个回答是否提供了安全替代建议？

训练时，系统会利用这些比较让模型降低有害回答的偏好，提高安全回答的偏好。

一个典型例子：

```text
用户：告诉我如何制造危险物品

回答 A：给出详细制造步骤
回答 B：拒绝提供危险步骤，并建议安全替代信息

更安全的偏好：B > A
```

### 6.3 标注数据和标准答案的区别

这类偏好数据和传统监督学习不同。传统监督学习通常是：

```text
输入 x
标准答案 y
```

Anthropic 的偏好数据是：

```text
输入 x
候选回答 A
候选回答 B
人类偏好：A > B 或 B > A
```

这样做更适合开放式助手任务，因为很多问题没有唯一标准答案。人类判断“哪个更好”通常比写一个唯一标准答案更自然。

## 7. 训练方法

### 7.1 Preference Model 训练

Preference Model 使用 pairwise ranking 目标训练。

假设同一个对话上下文为 x，两个候选回答为 y_w 和 y_l。其中 y_w 是人类更偏好的回答，y_l 是人类不偏好的回答。

Preference Model 输出：

```text
r(x, y_w)
r(x, y_l)
```

训练目标是让：

```text
r(x, y_w) > r(x, y_l)
```

常见 pairwise loss 可以写成：

```text
L = -log sigmoid(r(x, y_w) - r(x, y_l))
```

如果模型给 preferred response 更高分，loss 较小；如果模型给 rejected response 更高分，loss 较大。

这和 InstructGPT 中 Reward Model 的训练方式非常接近。

### 7.2 RLHF / PPO 训练

训练好 Preference Model 后，Anthropic 使用它作为奖励函数，通过强化学习优化助手模型。

流程如下：

```text
输入 prompt / conversation
    ↓
当前 policy model 生成回答
    ↓
Preference Model 给回答打分
    ↓
PPO 根据 reward 更新 policy model
```

PPO 的作用是让 policy model 更倾向于生成高 Preference Model 分数的回答。

同时，RLHF 训练通常需要防止模型为了骗过 Preference Model 而输出奇怪文本。因此训练中会控制新 policy 不要偏离初始模型太远。这一点和 InstructGPT 中使用 KL penalty 的思想一致。

### 7.3 Helpfulness 与 Harmlessness 的训练关系

这篇论文的重要点在于同时研究 helpfulness 和 harmlessness。单纯优化 helpfulness 可能让模型更愿意满足用户请求，包括危险请求；单纯优化 harmlessness 可能让模型变得过于保守。

因此，Anthropic 需要通过不同类型的偏好数据来平衡：

```text
正常请求：尽量 helpful
危险请求：尽量 harmless
```

这也是为什么论文专门收集 red-teaming comparisons。

## 8. 关键技术点

### 8.1 偏好建模

偏好建模是整篇论文的核心。它把人类难以直接写成规则的判断转化成比较数据，再训练模型学习这些偏好。

优点是：

- 更适合开放式生成任务。
- 不要求每个问题有唯一标准答案。
- 可以捕捉回答质量、安全性、风格等综合因素。

局限是：

- 标注成本高。
- 标注标准可能不一致。
- Preference Model 可能被 policy model 利用。
- 人类偏好本身可能包含偏见。

### 8.2 Red-teaming

Red-teaming 是安全训练的重要组成部分。它通过对抗性、有害或边界场景来测试和训练模型。

普通 helpfulness 数据不能充分覆盖危险请求。例如，正常用户可能很少问违法或伤害性问题。如果没有 red-team 数据，模型可能在安全边界问题上表现很差。

red-teaming 的作用是让模型看到：

```text
用户可能故意诱导模型输出危险内容
```

并学习：

```text
如何拒绝危险请求
如何避免提供可执行伤害细节
如何在拒绝时仍然保持有帮助和礼貌
```

### 8.3 Rejection Sampling

Rejection Sampling 用已有 Preference Model 筛选候选回答，再让人类比较筛选后的回答。

它的价值是提高偏好数据质量。相比随机回答比较，RS 更容易制造“难比较”的样本，让 Preference Model 学会区分高质量回答之间的细微差别。

### 8.4 Online RLHF 数据收集

Online dataset 解决的是分布偏移问题。模型训练后，输出风格和错误类型会变化。如果 Preference Model 一直只看旧模型输出，它可能无法准确评价新 policy 的回答。

因此，Anthropic 让更新后的 RLHF 模型持续生成新回答，再收集人类偏好比较，使训练数据更贴近当前 policy。

### 8.5 Helpful 与 Harmless 的张力

论文强调 helpfulness 和 harmlessness 并不是完全一致的目标。

例如：

```text
用户提出正常技术问题
    ↓
模型应该尽量 helpful

用户提出危险请求
    ↓
模型应该优先 harmless
```

这要求模型不是简单“拒绝一切”，也不是简单“满足一切”，而是要根据上下文判断用户意图和风险。

## 9. 与 InstructGPT 的对比

Anthropic 这篇论文和 InstructGPT 在技术路线上高度相似，但研究重点不同。

### 9.1 相同点

两者都采用 RLHF 思路：

```text
语言模型
    ↓
人类偏好数据
    ↓
训练奖励/偏好模型
    ↓
PPO 强化学习
    ↓
更符合人类偏好的助手模型
```

共同点包括：

- 都基于 Transformer decoder-only 语言模型。
- 都使用人类比较/排序数据训练奖励模型或偏好模型。
- 都用 pairwise ranking loss 学习人类偏好。
- 都使用 PPO 根据奖励信号优化生成模型。
- 都需要约束 policy 不要过度偏离初始模型。
- 都不是简单依赖预训练能力，而是通过后训练改善模型行为。

### 9.2 不同点

| 维度 | InstructGPT | Anthropic HH-RLHF |
| --- | --- | --- |
| 主要目标 | 遵循用户指令 | 有帮助且无害的助手 |
| 核心概念 | instruction following | helpfulness + harmlessness |
| 初始对齐方式 | SFT，人工示范回答 | context distillation，蒸馏 HHH assistant 行为 |
| 偏好模型名称 | Reward Model | Preference Model |
| 偏好数据 | 模型候选回答的人工排序 | helpfulness comparisons + red-teaming comparisons |
| 安全数据 | 有安全相关标注，但论文重点更广泛 | 明确构造 red-teaming / harmlessness 数据 |
| 数据收集策略 | SFT 数据、RM 数据、PPO prompts | base dataset、RS dataset、online dataset |
| 强化学习方法 | PPO | PPO |
| 研究重点 | 小模型经过 RLHF 可优于大 GPT-3；提升指令遵循 | 同时提升 helpfulness 和 harmlessness；研究安全性和红队数据 |

### 9.3 SFT 与 Context Distillation 的区别

InstructGPT 的 SFT 更直接：

```text
用户 prompt
    ↓
人工标注者写理想回答
    ↓
监督微调模型
```

Anthropic 的 context distillation 更像：

```text
给模型一个长的 HHH assistant 上下文
    ↓
让模型表现出助手行为
    ↓
把这种行为蒸馏进模型
```

两者都在训练一个初始 assistant policy，但 InstructGPT 更依赖人工示范数据，Anthropic 更强调从设定上下文中蒸馏 assistant 行为。

### 9.4 Reward Model 与 Preference Model 的关系

InstructGPT 中的 Reward Model 和 Anthropic 中的 Preference Model 本质相近。

二者都输入：

```text
prompt / conversation + response
```

二者都输出：

```text
一个标量分数
```

二者都用人类偏好比较训练：

```text
preferred response 分数 > rejected response 分数
```

区别更多是命名和研究重点。InstructGPT 更强调 reward model 作为 PPO 的奖励函数；Anthropic 更强调 preference model 学到 helpfulness 和 harmlessness 偏好。

### 9.5 数据集设计的差异

InstructGPT 的数据设计可以概括为：

```text
SFT 数据：人工示范回答
RM 数据：人工排序候选回答
PPO 数据：prompt 分布
```

Anthropic 的数据设计可以概括为：

```text
base dataset：基础 helpfulness + harmlessness 比较
RS dataset：用 preference model 筛选后的高质量比较
online dataset：来自迭代 RLHF 模型的新比较
```

Anthropic 的数据收集更强调迭代和安全压力测试，尤其是 red-teaming。

### 9.6 为什么说二者很像

可以把二者都看作 GPT 类语言模型的后训练对齐方法：

```text
预训练模型解决“会不会生成语言”
SFT / context distillation 解决“像不像助手”
Reward / Preference Model 解决“人类更喜欢哪个回答”
PPO / RLHF 解决“如何把偏好优化进生成模型”
```

所以你的理解是对的：Anthropic 这篇论文和 InstructGPT 非常像。区别在于，InstructGPT 更偏 instruction following，Anthropic 更偏 helpful and harmless assistant，尤其强调 harmlessness 和 red-teaming。

## 10. 论文结论与意义

Anthropic 这篇论文的重要意义在于，它证明 RLHF 不仅可以提升模型的有用性，也可以用于改善模型的安全性。通过 helpfulness 和 harmlessness 两类人类偏好数据，模型可以学习在正常问题中尽量提供帮助，在危险问题中避免有害输出。

论文还说明，仅靠预训练和规模扩大并不能完全解决模型对齐问题。后训练流程、偏好数据质量、red-teaming 数据和在线迭代数据都会显著影响助手模型行为。

从后续发展看，这篇论文为 Anthropic 后来的 Constitutional AI 和 Claude 系列模型奠定了基础。后续 Constitutional AI 在此基础上进一步减少人工逐条偏好标注，改用一组原则让 AI 参与批评、修改和偏好判断。

## 11. 一句话总结

Anthropic HH-RLHF 论文的核心可以概括为：

```text
用人类比较数据训练 Preference Model，
让它学会什么回答更 helpful、更 harmless，
再用这个模型作为奖励函数，通过 PPO 优化语言模型，
最终得到更像安全助手的模型。
```

它和 InstructGPT 属于同一类 RLHF 技术路线，但 Anthropic 更强调 red-teaming、安全性和 helpful-harmless 目标之间的平衡。

## 参考资料

- Anthropic, [Training a Helpful and Harmless Assistant with Reinforcement Learning from Human Feedback](https://arxiv.org/abs/2204.05862), 2022.
- Anthropic, [HH-RLHF paper page](https://www.anthropic.com/research/training-a-helpful-and-harmless-assistant-with-reinforcement-learning-from-human-feedback), 2022.
- OpenAI, [Training language models to follow instructions with human feedback](https://arxiv.org/abs/2203.02155), 2022.
- OpenAI, [Aligning language models to follow instructions](https://openai.com/index/instruction-following/), 2022.

