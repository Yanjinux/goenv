<!--
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-14 10:58:54
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-14 11:01:39
 * @FilePath: \goenv\code\doc\auth.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# 权限设计


## 常用模式

### RBAC
1. RBAC0
用户-》角色-》权限


#### 权限
1. 页面权限
2. 操作权限
3. 数据权限
#### 校色


2. RBAC1 
具有继承关系


3. RBAC2 
添加了责任分离校色互斥

4. RBAC3 
结合 1和2




### 设计

```
用户是怎么分类的（用户角色）
用户和用户之间是否有关系？如果有，是什么关系？关系是什么结构的？
如公司组织架构那种层级分明的树形结
```

1. 用户分组