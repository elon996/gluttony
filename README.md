# Gluttony

Gluttony是一个http cve poc框架，使用它可以快速检测内网cve。

#### 1.安装

```
go get github.com/elon996/gluttony
```



#### 2.用法

检测192.168.1.1的thinkphp漏洞

p参数可以是文件或者文件夹

```
gluttony scan scan -u 192.168.1.1 -p cve-2018-10225.yaml
gluttony scan scan -u 192.168.1.1 -p cve
```



#### 3.poc编写

poc应由yaml编写，标准模板如下：

```yaml
single: true
cve:
  id: cve-2021-0222
  product: thinkphp

requests:
  - method: get
    redirect: false
    path:
      - /{{.name}}
    headers:
      - User-Agent: 
    body: asd
    detections:
      - res.code==200
    set:
      - name: asd

```

##### 1.single

single为连续检测标志，使用后requests一次只发出一个http包，并且前一个http包的detections若不为match，则检测失败，不再发送后一个http包

##### 2.multiple

默认发送属性为multiple，即连续发送requests中的http包

##### 3.cve

cve为cve信息，方便查询

##### 4.requests

requests是一组http包，能直接配置http包的信息

- method为http 方法

- redirect为重定向，可以设置重定向次数,默认20次

  ```yaml
  redirect: 1
  ```

  

- path是一组路径

  - path可以使用set中设置的值，示例如下：

  - ```yaml
    path:
     - /{{.name}}
    set:
     - name: asd
    ```

  

- body是post内容，默认为x-www-form-urlencoded，，示例如下：

  - ```yaml
    body: user=qwe
    ```

    

  支持xml和json，但需要在headers中设置，示例如下：

  - ```
    headers:
     - Content-Type: application/json
    body: |
     {"asd":"qwe"}
    ```

- authuser和authpassword是http基本认证

- proxy为http代理，示例如下：

  - ```
    proxy: 
     - http://127.0.0.1:8080
    ```

    

  

##### 5.detections

detections为js语法，目前支持res对象：

```javascript
res.code == 200  //返回true
res.body.search("body")   //查找response的body的"body"字符串,存在返回true
res.length <= 100
res.header.search("Content-Type")  //查找response的header的"Content-Type"字符串,存在返回true
```

detections编写时只应该返回布尔值和数字，否则报错

**注意：勿使用他人poc，可能引发安全问题**

使用它们能编写各式各样poc



#### 4.其他用法

1.用gluttony定制webshell

在某些场景中，大众化的webshell是不可用的，编写gluttony poc可以解决特定场景的webshell执行



