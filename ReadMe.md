# 通过OpenCV中的查找轮廓批量抠图

## 示例一 抠出下图中的身份证:



![demo.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lype3czej303r050q3j.jpg)

step1:

![QQ截图20221129142246.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lz807343j30aj08u762.jpg)

step2:

![QQ截图20221129142331.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lz8tus92j30ag09cdgn.jpg)

step3:

![QQ截图20221129142406.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lz9e00gaj30fs098q50.jpg)

step4:

![QQ截图20221129142457.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzaf9vhqj30fm08a0v2.jpg)

step5:

![QQ截图20221129142709.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzcu9lhaj30i106dabm.jpg)

step6:

![QQ截图20221129142803.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzdowqk4j30kw03s3zj.jpg)

批量操作:

![QQ截图20221129142952.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzgh2o7lj30tg0emjxz.jpg)

## 示例二 抠出下图中的身份证中的人像:



![demo2.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lypyb5glj303001pgll.jpg)



step3:中 `ThresholdBinary` 修改为 `ThresholdBinaryInv`

![QQ截图20221129143313.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzjhdpvfj30gt05xdhi.jpg)

step4:

![QQ截图20221129143329.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzjy9d4ej30f605fq4a.jpg)

## 示例三 抠出下图中的人脸



![demo3.jpg](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzkink8tj30xc0idq4o.jpg)

### 通过人脸定位:

step3:

![QQ截图20221129143740.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzp7sp36j30ld05ntau.jpg)

step4: 选择第三大的人脸

![QQ截图20221129143803.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzpo5r6bj30rc0bcn3q.jpg)

step5:

![QQ截图20221129143910.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8lzpqxs2yj30t309mdkp.jpg)

通过头发定位:

step3:

![QQ截图20221129145021.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8m02yzypzj30ne05stak.jpg)

step5:

![QQ截图20221129145034.png](https://tva1.sinaimg.cn/large/006ulzy2ly1h8m0383oh6j30nd06mgog.jpg)