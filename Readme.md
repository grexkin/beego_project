# beego Study
- 2020-0902: 1~8
- 2020-0903: 9~14
- 2020-0904: 14~20
- 2020-0906: 20~25
- 2020-09-07: 26~29
- 2020-09-08：29~33
- 2020-09-10: 34~36
## 数据交互小结
1. 后端数据传输到前端展示
    - 字符串渲染
    - 数组
    - 结构体
    - 结构体+数组
    - map渲染
    - map+结构体渲染
    - 切片
2. 静态文件的使用
    - 修改默认静态文件路径
3. 前端传数据到后端
    - get
        - /xxx?name=xxx
        - /xxx/111
    - post
        - action： URL
        - method: post
    - ajax
        - button
        - jquery $.ajax({url:xxx,type:xxx,data: JSON,dataType：xxx,success:xxx,error:xxx})
            - url
            - type
            - data
            - dataType
            - 回调
                - success
                - error
4. 后端传输数据到前端
    - json
    - jsonp
    - xml
    - yaml
    > key 要与上述对应

5. flash
    数据校验
    - Notice
    - Error
    - Store
    - redirct
    - 对ajax提交的数据进行校验
