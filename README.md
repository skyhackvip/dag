# 分布式任务调度系统之任务编排及工作流实现原理与 golang 实践

定时调度系统（定时任务、定时执行）算是工作中经常依赖的中间件系统，简单使用操作系统的 crontab，或基于 Quartz，xxl-job 来搭建任务调度平台，行业有很多优秀的开源产品和中间件。了解其工作和设计原理，有助于我们完善或定制一套适合公司业务场景的任务调度中间件。

今天我们对调度任务的依赖关系及编排展开分析，实现一套工作流，来满足任务间的复杂依赖的场景。

本章内容提要：
- 任务调度依赖 & 工作流
- 图相关知识
- golang 并发相关

### 阅读全文链接
[任务编排及工作流实现](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483920&idx=1&sn=922d3d3e6ff07951fd45d629a960dca3&chksm=e8215d00df56d41679227c299134e7cea9cb9f93094f8597abfd3675a19116da628e02e28042&token=1776095493&lang=zh_CN#rd)

扫码关注微信订阅号支持：

![技术岁月](https://raw.githubusercontent.com/skyhackvip/ratelimit/master/techyears.jpg)
