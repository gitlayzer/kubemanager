# kubemanager-api
这是kubemanager的后端API程序，完全由Go语言编写，目前还在开发中，部分功能还不太完善，可以拿走二次开发。

本次版本属于二次迭代版本，后续若更新会以小版本形式发布，本次更新内容如下：
```shell
1：本次版本为API程序添加了启动参数，具体参数可以使用-h列出，并且可以在启动API服务时指定参数，增加了更为灵活的自定义性
2：本次更新参数如下所示：
```

```shell
  -kubeconfig string                              
        Set kubeconfig file (default "kubeconfig")
  -listen string                                  
        Set address (default "0.0.0.0:9090")      
  -pass string                                    
        Set password (default "123456")           
  -podlogtaillines int                            
        Set pod log tail lines (default 2000)     
  -user string                                    
        Set admin (default "admin")
```
