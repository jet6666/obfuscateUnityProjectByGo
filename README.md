>本工具是为上架ios马甲包而准备的, 请参考Unity C#代码级混淆方案   **[Unity游戏iOS马甲包处理方案](https://zhuanlan.zhihu.com/p/523090660)**

 

## 1 GO代码执行环境

- go mod （go mod tidy)
- 单个文件执行
 

## 2 执行顺序
- 请参考代码中注释选择路径
-  csharp_add_spam.go ，将原cs代码加入垃圾代码，可不执行

- prefab_comment.go 将*.prefab文件加入垃圾代码，可不执行

- csharp_rename_step1.go 关掉unity,将cs 类改名
- 上一步执行正确后，打开unity,等待编译完成
-csharp_rename_step2.go 完成改名


- prefab_confuse_step1.go 关掉unity,将prefab类改名
- 上一步执行正确后，打开unity,等待编译完成
-prefab_confuse_step2.go 完成改名