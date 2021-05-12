# notify
解决 gitlab ci 无法主动发送邮件的问题

# 安装
```
go get github.com/hellojukay/notify
```
gitlab pipeline 设置如下
```yaml
stages:
  - notify
  
send email:
  stage: notify
  image: hellojukay/email:1.0.1
  script:
    # 发送邮件的代码参考 https://github.com/hellojukay/notify
    - notify -smtp-port=587 -smtp-server=mail.xxx.com -user=$EMAIL_USER  -to=$RECEIVER -smtp-pass=$EMAAL_PASS -path=email.html
```
