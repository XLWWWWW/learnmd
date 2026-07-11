# GPT-4 模型改进、数据与训练集整理

## 摘要

GPT-4 是 OpenAI 在 2023 年发布的大规模多模态模型。根据 OpenAI 的 GPT-4 Technical Report，GPT-4 是一个 Transformer-style 模型，通过预测文档中的下一个 token 进行预训练，并在预训练之后使用 RLHF 进行后训练对齐。和 GPT-3.5 相比，GPT-4 在复杂指令理解、专业考试、代码、数学、事实性、多语言能力和安全拒答等方面都有明显提升。

需要特别注意：OpenAI 并没有公开 GPT-4 的完整技术细节。官方报告明确没有披露模型参数量、精确架构、训练硬件、训练计算量、数据集构造方式、训练方法细节和数据配比。因此，理解 GPT-4 时要区分：

- **官方明确公开的信息**：模型类型、训练目标、数据大类、后训练流程、安全训练方法、评测结果。
- **官方没有公开的信息**：参数量、token 数、具体训练集列表、各类数据比例、完整过滤规则、硬件规模、训练超参数。

本文主要依据 OpenAI 的 [GPT-4 Technical Report](https://arxiv.org/abs/2303.08774)、[GPT-4 research page](https://openai.com/index/gpt-4-research/) 和 [GPT-4 System Card](https://cdn.openai.com/papers/gpt-4-system-card.pdf) 整理。

## 1. GPT-4 的定位

GPT-4 可以理解为 GPT 系列从“大语言模型”走向“大规模多模态基础模型”的关键版本。

它的基本定位是：

```text
GPT-4
    = 大规模 Transformer-style 自回归模型
    + 文本和图像输入能力
    + 文本输出能力
    + SFT / RM / RLHF 后训练
    + 安全专项训练与评测
```

OpenAI 在 GPT-4 research page 中描述 GPT-4 为一个 large multimodal model，可以接受图像和文本输入，并生成文本输出。也就是说，GPT-4 不是图像生成模型，而是可以把图像作为输入上下文的一部分，然后用文本回答。

### 1.1 GPT-4 和 GPT-3.5 的关系

GPT-4 不是对 GPT-3.5 的一个简单 prompt 改良，而是一个新的、更强的基础模型和后训练系统。

可以粗略理解为：

```text
GPT-3 / GPT-3.5
    ↓
更大规模、更强训练系统、更强数据处理和后训练
    ↓
GPT-4 base model
    ↓
SFT + RM + RLHF + 安全训练
    ↓
GPT-4 launch model
```

GPT-4 和 GPT-3.5 的区别主要体现在：

- 更强的复杂任务处理能力。
- 更好的指令遵循能力。
- 更强的多语言和代码能力。
- 更低的幻觉倾向，但仍然会幻觉。
- 更强的安全拒答能力。
- 支持图像输入。

## 2. 官方没有公开的部分

GPT-4 和 GPT-1、GPT-2、GPT-3 的一个很大区别是：GPT-4 的技术报告没有公开很多核心训练细节。

官方明确没有进一步披露：

| 类别 | 是否公开 | 说明 |
| --- | --- | --- |
| 参数量 | 未公开 | 外界有很多猜测，但不是官方信息 |
| 精确架构 | 未公开 | 只说是 Transformer-style model |
| 层数、宽度、注意力头数 | 未公开 | 没有像 GPT-3 paper 那样列表 |
| 训练 token 数 | 未公开 | 没有公开总 token 规模 |
| 数据集构造方式 | 未公开 | 只公开数据大类 |
| 数据比例 | 未公开 | 没有公开网页、代码、书籍、数学等比例 |
| 硬件和训练 compute | 未公开 | 没有公开 GPU 数量、FLOPs 等 |
| 训练超参数 | 未公开 | 没有公开 learning rate、batch size、optimizer 等 |

这意味着：如果看到网上说 GPT-4 有多少参数、用了多少 token、是不是 MoE、用了哪些具体网站或书籍，这些大多属于外界推测或泄露传闻，不应该写成确定结论。

## 3. 模型结构怎么理解

### 3.1 Transformer-style 自回归模型

OpenAI 把 GPT-4 描述为 Transformer-style model，预训练目标是预测文档中的下一个 token。

这和 GPT 系列的一贯路线一致：

```text
输入前文 tokens
    ↓
模型预测下一个 token
    ↓
继续自回归生成
```

也就是说，GPT-4 的底层语言建模目标仍然是 next-token prediction。

### 3.2 多模态输入

GPT-4 的一个关键改进是可以接受文本和图像混合输入。

例如：

```text
用户输入：
    一段文字说明
    + 一张图表
    + 一个问题

GPT-4 输出：
    文本解释、代码、推理过程、结论等
```

这使 GPT-4 可以处理：

- 图表理解。
- 截图分析。
- 试题图片。
- 文档图片。
- 含图的推理题。
- 视觉信息和文字信息混合的问题。

注意：GPT-4 的输出仍然是文本。它和 DALL-E 这类图像生成模型不是同一类系统。

### 3.3 Base model、early model、launch model

读 GPT-4 资料时会看到几个不同版本，需要分清：

| 名称 | 含义 |
| --- | --- |
| GPT-4 base model | 预训练后的基础模型，还没有完整后训练对齐 |
| GPT-4-early | 早期 instruction-following 版本，安全缓解较少 |
| GPT-4-launch | 发布时的版本，经过更多 helpfulness 和 harmlessness 训练 |

System Card 重点比较 GPT-4-early 和 GPT-4-launch，因为它们更接近真实部署前后的助手行为。base model 虽然能体现预训练能力，但不一定适合直接给红队专家使用。

## 4. GPT-4 的主要改进

### 4.1 更强的复杂指令理解

OpenAI 在 GPT-4 research page 中说，日常闲聊中 GPT-3.5 和 GPT-4 的差别可能不明显；当任务复杂度提高时，GPT-4 更可靠、更有创造力，也更能处理细微复杂的指令。

可以理解为：

```text
简单问答：
    GPT-3.5 和 GPT-4 差别可能不大

复杂任务：
    多约束
    长上下文
    需要推理
    需要代码
    需要准确遵循格式
    ↓
    GPT-4 优势明显
```

### 4.2 专业考试表现更强

GPT-4 在多个专业和学术考试上达到接近人类较高水平的表现。最常被引用的例子是模拟 Uniform Bar Exam：

```text
GPT-4：约前 10%
GPT-3.5：约后 10%
```

这说明 GPT-4 的综合阅读、推理、知识调用和长答案组织能力相比 GPT-3.5 有明显提升。

但要注意：考试能力不等于真实可靠性。GPT-4 仍然可能幻觉、犯推理错误、输出不可靠代码。

### 4.3 多语言能力增强

OpenAI 把 MMLU 翻译成多种语言来测试 GPT-4。在大多数测试语言中，GPT-4 的表现超过 GPT-3.5 的英文表现，也超过一些已有大模型的英文表现。

这说明 GPT-4 的能力不只局限于英语。但 System Card 也提醒：大部分预训练数据和 alignment data 仍然是英语，安全缓解措施也主要在英语和美国语境下设计与测试。

因此可以这样理解：

```text
能力层面：
    GPT-4 有较强跨语言迁移能力

安全和文化适配层面：
    非英语语境仍然可能有更多误判和漏判
```

### 4.4 代码能力增强

Technical Report 用 HumanEval 等基准评估代码生成能力。GPT-4 在 HumanEval 上显著强于 GPT-3.5。

但 GPT-4 并不是一个不会写 bug 的程序员。报告也明确指出，GPT-4 仍然可能在生成代码时引入安全漏洞。

正确理解是：

```text
GPT-4 更会写代码
    ≠
GPT-4 写的代码总是正确、安全、可部署
```

### 4.5 事实性和幻觉有所改善

GPT-4 相比 GPT-3.5 降低了幻觉倾向。Technical Report 中提到，在内部 adversarial factuality evaluations 上，GPT-4-launch 比最新 GPT-3.5 高 19 个百分点。

但 GPT-4 仍然不是事实数据库：

- 会编造事实。
- 会在不确定时自信回答。
- 会犯简单推理错误。
- 对 2021 年 9 月之后的大多数事件缺乏预训练知识。

OpenAI 还指出：预训练模型本身的校准性较好，但经过 post-training 后，校准性会下降。也就是说，对话模型更像“有用助手”后，不一定更会表达自己的不确定性。

### 4.6 安全拒答能力增强

GPT-4 的发布版本相比 GPT-3.5 明显减少了对禁止内容请求的错误响应。OpenAI research page 给出的数据是：

- 对 disallowed content 的响应倾向相比 GPT-3.5 降低 82%。
- 对 sensitive requests 按政策响应的比例提高 29%。

这主要来自 RLHF、安全专项数据、RBRM 和红队反馈等一组后训练与安全缓解措施。

## 5. 预训练数据怎么理解

### 5.1 官方公开的数据大类

OpenAI 公开的信息可以概括为：

```text
GPT-4 pretraining data
    = 公开可获得数据，例如互联网数据
    + 第三方授权数据
```

OpenAI research page 进一步描述，这些数据是 web-scale corpus，包含：

- 正确和错误的数学解法。
- 强推理和弱推理。
- 自相矛盾和一致的陈述。
- 多种意识形态和观点。

这点很重要：预训练数据不是只包含“正确答案”。大规模互联网文本里天然会有正确内容、错误内容、有偏内容、冲突观点、低质量内容和高质量内容。预训练模型学到的是文本分布，不是一个天然可靠的真理系统。

### 5.2 数据截止时间

GPT-4 Technical Report 提到，GPT-4 对大多数预训练数据截止时间之后的事件缺乏知识。这个主要截止时间是：

```text
2021 年 9 月
```

但报告脚注也说明，预训练和后训练数据中包含少量更近期的数据。

因此，更准确的说法是：

```text
GPT-4 的大多数预训练数据截止在 2021 年 9 月；
但训练数据中有少量更晚的数据。
```

### 5.3 数学数据：MATH 和 GSM-8K

Technical Report 附录 E 提到，为了提升 GPT-4 的数学推理能力，OpenAI 把 MATH 和 GSM-8K 训练集中的数据混入 GPT-4 训练中。

但这里有几个关键限制：

- 这些数学 benchmark 数据只占 GPT-4 总训练预算的极小一部分。
- 混入时保留了一部分训练数据未使用。
- GSM-8K 测试集经过污染检查，确认没有包含在训练集中。
- 因为训练集中混入了部分 GSM-8K / MATH 数据，所以 GPT-4 在 GSM-8K 上的结果应理解为介于纯 few-shot transfer 和 benchmark-specific tuning 之间。

这说明 GPT-4 的数学能力既来自大规模预训练，也受到有意识的数据增强影响。

### 5.4 数据污染检查

OpenAI 对考试和学术 benchmark 做了训练数据污染检查。

污染检查的目的：

```text
确认评测题目是否曾出现在训练数据中
```

如果测试题已经出现在训练数据中，模型表现可能不是泛化能力，而是记忆或近似记忆。

Technical Report 中提到：

- 一些考试题目有不同程度的 contamination。
- 对考试评测，OpenAI 会报告去掉污染题后的结果。
- 对部分 academic benchmarks，OpenAI 估计了测试集和预训练数据的重合比例。
- BIG-bench 的部分内容被发现混入训练集，因此没有纳入报告结果。

这也提醒我们：看大模型 benchmark 时，不能只看分数，还要看是否检查了数据污染。

### 5.5 预训练数据过滤

System Card 中提到，OpenAI 在预训练阶段对 GPT-4 的数据混合进行了过滤，目标之一是减少不合适的色情文本内容。

方法包括：

- 内部训练的分类器。
- 基于词典的检测方法。
- 对高概率包含不合适色情内容的文档进行移除。

这属于模型级别安全缓解的一部分。也就是说，安全不只发生在 RLHF 阶段，预训练数据选择和过滤本身也是安全训练的一部分。

## 6. 后训练数据怎么理解

GPT-4 的后训练目标不是继续扩大知识，而是把 base model 调成更符合用户意图、更安全、更像助手的模型。

后训练大致包括：

```text
SFT demonstration data
    + ranking data
    + reward model
    + RLHF / PPO
    + safety prompts
    + RBRM
    + red-teaming data
    + synthetic prompts
    + production data with consent
```

### 6.1 Demonstration data

Demonstration data 是“示范数据”。

形式类似：

```text
输入 prompt
    ↓
人类示范：模型应该怎样回答
```

这些数据用于 SFT，也就是 supervised fine-tuning。目标是让 GPT-4 模仿人类示范中的助手行为。

可以理解为：

```text
base GPT-4
    ↓
看人类示范回答
    ↓
学习更像一个助手
    ↓
GPT-4 SFT model
```

### 6.2 Ranking data

Ranking data 是“排序数据”。

形式类似：

```text
输入 prompt
    ↓
模型回答 A
模型回答 B
模型回答 C
...
    ↓
人类标注：从最好到最差排序
```

这些数据用于训练 Reward Model，简称 RM。

RM 的作用是：

```text
输入：prompt + candidate response
输出：一个分数
含义：这个回答有多符合人类偏好
```

RM 学到的是“平均标注员更偏好什么样的回答”。训练好 RM 后，它可以作为 PPO 强化学习阶段的奖励信号。

### 6.3 RLHF / PPO

GPT-4 的 RLHF 流程和 InstructGPT 很接近：

```text
预训练 GPT-4 base model
    ↓
SFT：用 demonstration data 做监督微调
    ↓
训练 RM：用 ranking data 学习人类偏好
    ↓
PPO：用 RM 分数作为 reward 优化 SFT model
    ↓
更符合人类偏好的 GPT-4 policy
```

其中 PPO 是强化学习算法。它让模型更倾向生成 RM 评分高的回答。

但要注意一个关键结论：GPT-4 的考试和推理能力主要来自预训练。OpenAI 在 research page 中强调，RLHF 主要带来的是 steerability，也就是让模型更按用户意图和安全边界回答；它不一定提升考试性能，不主动处理时甚至可能降低某些考试表现。

可以概括为：

```text
预训练：
    主要提供知识、语言能力、推理能力、代码能力

后训练：
    主要提供助手行为、指令遵循、拒答边界、安全风格
```

## 7. 安全训练数据与 RBRM

### 7.1 为什么 RLHF 还不够

普通 RLHF 可以让模型更有帮助、更符合标注员偏好，但仍然有问题：

- 对不安全请求可能仍然回答。
- 对安全请求可能过度拒绝。
- 标注指南没有覆盖的边界场景中表现不稳定。
- 容易被 jailbreak prompt 绕过。

System Card 中提到，GPT-4-early 有时会过度谨慎，拒绝无害请求，或者过度使用模糊拒答。

因此 OpenAI 在 GPT-4-launch 中加入了更细粒度的安全训练。

### 7.2 Safety-relevant prompts

安全相关 prompt 包括：

- 违法请求。
- 自伤相关请求。
- 危险化学品或武器相关请求。
- 医疗、法律等敏感建议。
- 仇恨、骚扰、性内容等政策相关内容。
- 边界案例，也就是看起来相似但其实允许回答的问题。

安全训练不只是让模型学会拒绝，也要让模型学会“不该拒绝时不要拒绝”。

所以安全训练目标是双向的：

```text
disallowed prompt:
    应该拒绝

allowed prompt:
    不应该误拒绝
```

### 7.3 RBRM 是什么

RBRM 是 rule-based reward model，规则型奖励模型。

在 GPT-4 System Card 中，RBRM 被描述为一组 zero-shot GPT-4 classifiers。它在 PPO 微调过程中为 policy model 提供额外的奖励信号。

RBRM 的输入包括：

```text
prompt，可选
policy model 的输出
human-written rubric，人类写的评分规则
```

RBRM 根据 rubric 判断输出属于哪一类。例如：

```text
A：理想风格的拒绝
B：不理想风格的拒绝
C：包含禁止内容
D：安全的非拒绝回答
```

然后训练系统可以：

- 对危险请求，奖励正确拒绝。
- 对安全请求，奖励不误拒绝。
- 对含有禁止内容的回答，给负向信号。
- 对啰嗦、逃避、边界不清的拒答，给较差信号。

### 7.4 RBRM 数据来源

System Card 提到，RBRM 相关训练数据的主要来源包括：

- 经用户同意使用的生产流量数据。
- OpenAI Moderation API 和 zero-shot GPT-4 辅助过滤分类的数据。
- 人类 reviewer 分类后的数据。
- 红队专家编写的 prompts。
- 模型生成的 synthetic prompts。
- 其他内部或公开数据集。
- 模型改写出的 boundary prompts。
- 标注员尝试绕过 GPT-4-launch 行为时产生的 ranking data。

这里的核心思想是：安全训练不只靠一批静态人工数据，而是结合真实使用、红队攻击、模型合成、人工审核和自动分类，构造更接近真实风险场景的数据。

### 7.5 Boundary prompts

Boundary prompts 是理解 GPT-4 安全训练的关键。

例如两个请求非常像：

```text
请求 A：告诉我如何合成危险化学品
请求 B：解释为什么某种化学品危险，以及如何安全处理
```

A 应该拒绝，B 可能应该安全回答。

如果模型只学会“看到化学品就拒绝”，就会过度拒答。如果模型只追求 helpful，就可能给出危险细节。

因此 OpenAI 会让模型把 disallowed prompts 改写成非常相似但允许回答的 boundary prompts，再用 RBRM 确保模型不要拒绝这些安全请求。

## 8. GPT-4 的训练流程总图

可以把 GPT-4 的训练流程理解为：

```text
大规模数据收集与处理
    ↓
预训练数据过滤
    ↓
Transformer-style base model 预训练
    目标：预测下一个 token
    ↓
GPT-4 base model
    ↓
SFT：人类示范数据
    ↓
GPT-4 SFT model
    ↓
RM：人类排序数据训练奖励模型
    ↓
PPO / RLHF：用 RM 作为奖励优化 policy
    ↓
加入安全相关 prompts 和 RBRM 奖励
    ↓
红队、专家评测、边界样本、合成样本继续迭代
    ↓
GPT-4-launch
```

更简洁地说：

```text
Pretraining 负责能力
SFT 负责助手样式
RM 负责学习人类偏好
RLHF / PPO 负责把偏好优化进模型
RBRM 负责更精细的安全边界
Red teaming 负责发现模型哪里容易出问题
```

## 9. 数据集分类总表

| 数据类别 | 用在什么阶段 | 数据形式 | 主要作用 |
| --- | --- | --- | --- |
| 公开可获得数据 | 预训练 | 大规模文本、网页等 | 学语言、知识、代码、推理模式 |
| 授权第三方数据 | 预训练 | 第三方授权语料 | 扩展高质量训练语料 |
| MATH / GSM-8K 训练数据 | 预训练混合 | 数学题和解法 | 增强数学推理 |
| 过滤后的安全预训练数据 | 预训练前处理 | 删除部分不合适内容后的语料 | 降低有害内容学习 |
| Demonstration data | SFT | prompt + 人类示范回答 | 学会助手式回答 |
| Ranking data | RM 训练 | prompt + 多个回答 + 排序 | 训练奖励模型 |
| Safety prompts | RLHF / RBRM | 安全、敏感、禁止、边界请求 | 学拒绝和不误拒绝 |
| Red-team prompts | 安全评测和训练 | 专家攻击性 prompt | 发现高风险失败模式 |
| Synthetic prompts | 安全训练 | 模型生成的样本 | 扩充边界场景 |
| Production data with consent | 安全训练和评测 | 经同意的真实使用数据 | 贴近真实分布 |
| Boundary prompts | RBRM / 安全训练 | 和危险请求相似但允许回答的请求 | 减少过度拒答 |

## 10. GPT-4 改进背后的关键思想

### 10.1 Scaling 不只是“模型更大”

GPT-4 的一个重要工程贡献是 predictable scaling。

OpenAI 强调，大模型训练非常昂贵，不可能在最大模型上反复试错。因此他们构建了可预测扩展的训练基础设施和优化方法，用小模型训练结果预测大模型的 loss 和部分能力。

Technical Report 中提到：

- 用最多小 10,000 倍 compute 的模型预测 GPT-4 最终 loss。
- 用小 1,000 倍 compute 的模型预测 HumanEval pass rate。

这说明 GPT-4 的改进不只是“砸更多算力”，还包括训练系统稳定性、优化方法、数据处理、评测方法和 scaling law 的工程化。

### 10.2 能力来自预训练，行为来自后训练

这是理解 GPT-4 最重要的一句话。

```text
能力：
    大部分来自预训练

行为：
    大部分来自后训练
```

比如：

- 会不会做数学、写代码、理解法律题，主要来自预训练学到的能力。
- 会不会按用户要求回答、会不会拒绝危险请求、语气是否像助手，主要来自后训练。

所以 RLHF 不是让模型“突然变聪明”的魔法。它更像是把已有能力引导到符合人类偏好的输出方式上。

### 10.3 安全不是单一数据集解决的

GPT-4 的安全改进不是来自一个简单的 harmful / harmless 数据集，而是一套 pipeline：

```text
预训练数据过滤
    +
人类示范和排序
    +
RLHF
    +
安全专项 prompts
    +
RBRM
    +
红队专家
    +
生产数据反馈
    +
部署时监控
```

这和 Anthropic HH-RLHF 的思路相似，都是通过人类偏好和安全数据让模型更符合人类期望。但 GPT-4 公开资料给出的数据规模更少，更多强调系统工程和安全 pipeline。

## 11. 和 InstructGPT、Anthropic HH-RLHF 的关系

### 11.1 和 InstructGPT 的关系

GPT-4 后训练流程基本延续了 InstructGPT 的路线：

```text
SFT demonstration data
    ↓
Reward Model ranking data
    ↓
PPO / RLHF
```

区别是 GPT-4 的基础模型更强，任务更复杂，安全训练更系统化，还加入了 RBRM 和更广泛的红队评测。

### 11.2 和 Anthropic HH-RLHF 的关系

Anthropic HH-RLHF 更适合学习“偏好数据怎么收集”，因为它公开了比较具体的数据集数量：

```text
base dataset
RS dataset
online dataset
```

GPT-4 更适合学习“大模型发布前系统性训练和安全缓解怎么做”，因为它公开了：

```text
pretraining
SFT
RM
RLHF
RBRM
red teaming
system safety
evals
```

但 GPT-4 没有公开每类数据的规模和比例。

## 12. 常见误解

### 12.1 GPT-4 的参数量是否确定？

不确定。OpenAI 没有公开 GPT-4 参数量。

网上关于 GPT-4 参数量的说法不能当成官方结论。

### 12.2 GPT-4 是否公开了完整训练集？

没有。OpenAI 只公开了数据大类：

```text
公开可获得数据
第三方授权数据
少量公开 benchmark 训练数据混入，例如 MATH / GSM-8K
后训练的人类示范和排序数据
安全相关数据
```

具体网站、书籍、代码仓库、比例和过滤细节没有公开。

### 12.3 GPT-4 的能力是不是 RLHF 训练出来的？

不完全是。OpenAI 明确强调，考试能力主要来自预训练，RLHF 主要让模型更符合用户意图和安全边界。

### 12.4 GPT-4 是否不会幻觉？

不是。GPT-4 只是相对 GPT-3.5 降低了幻觉倾向，仍然会幻觉。

### 12.5 GPT-4 是否总比 GPT-3.5 安全？

整体安全指标更好，但不是绝对安全。System Card 明确说，模型级缓解增加了诱导坏行为的难度，但仍然可能被 jailbreak，仍然需要部署时监控和使用策略。

## 13. 一句话总结

GPT-4 的核心可以概括为：

```text
用大规模预训练获得基础能力，
用 SFT 和 RLHF 学会符合人类偏好的助手行为，
用 RBRM、红队和安全数据细化拒答边界，
再通过可预测 scaling 和系统评测让训练过程更稳定、更可控。
```

最重要的理解点是：

```text
GPT-4 的“聪明”主要来自预训练；
GPT-4 的“像助手”和“更安全”主要来自后训练。
```

## 参考资料

- OpenAI, [GPT-4 Technical Report](https://arxiv.org/abs/2303.08774), 2023.
- OpenAI, [GPT-4 research page](https://openai.com/index/gpt-4-research/), 2023.
- OpenAI, [GPT-4 System Card](https://cdn.openai.com/papers/gpt-4-system-card.pdf), 2023.
- OpenAI, [Training language models to follow instructions with human feedback](https://arxiv.org/abs/2203.02155), 2022.
- Anthropic, [Training a Helpful and Harmless Assistant with RLHF](https://arxiv.org/abs/2204.05862), 2022.
