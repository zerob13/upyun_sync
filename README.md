# goDashboard-for-UPYUN
##介绍    
* 基于upyun_go_SDK 增加了一些SDK接口方便开发使用    
* 前端基于Bootstrap-3-Admin-Theme 开发，样式漂亮移动端兼容性好    
* 主要作用是用来管理upyun的空间中的各种文件，具有基本的上传，删除等功能，后续还会增加各种实用功能    
* 尽量做到操作简单，让不懂技术的小白也能够管理好自己的云    

##部署    
```
cd $GOPATH/src
go get github.com/astaxie/beego
go get github.com/beego/bee
git clone git@gitcafe.com:zerob13/goDashboard-for-UPYUN.git upyun_sync
cd upyun_sync
bee run

```
##Demo (并非基于最新代码)       
[http://zerob13.in:6699](http://zerob13.in:6699)
