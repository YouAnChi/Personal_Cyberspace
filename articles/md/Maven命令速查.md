# Maven命令速查

Maven是一个强大的项目管理和构建工具，广泛用于Java项目中。以下是按分类整理的Maven常用命令的详细表格。

| 分类         | 命令                                      | 描述                                                 |
| :----------- | :---------------------------------------- | :--------------------------------------------------- |
| **项目管理** | `mvn archetype:generate`                  | 根据模板创建一个新的Maven项目                        |
|              | `mvn clean`                               | 删除项目生成的所有文件（target目录）                 |
|              | `mvn validate`                            | 验证项目是否正确且所有必要信息都可用                 |
|              | `mvn compile`                             | 编译项目的源代码                                     |
|              | `mvn test-compile`                        | 编译测试源代码                                       |
|              | `mvn test`                                | 运行测试                                             |
|              | `mvn package`                             | 编译代码并打包成可分发格式（如JAR、WAR）             |
|              | `mvn verify`                              | 运行任何检查以验证包是有效的且满足质量标准           |
|              | `mvn install`                             | 将包安装到本地仓库，这样在本机的其他项目中就可以使用 |
|              | `mvn deploy`                              | 将最终的包复制到远程仓库以共享给其他开发人员和项目   |
| **依赖管理** | `mvn dependency:resolve`                  | 解析项目的所有依赖                                   |
|              | `mvn dependency:tree`                     | 显示项目依赖树                                       |
|              | `mvn dependency:analyze`                  | 分析项目的依赖使用情况，报告未使用的依赖             |
| **插件管理** | `mvn help:describe -Dplugin=plugin`       | 显示插件的详细信息                                   |
|              | `mvn help:all-profiles`                   | 显示项目中所有可用的配置文件                         |
| **运行项目** | `mvn exec:java`                           | 运行一个指定的类                                     |
|              | `mvn exec:exec`                           | 运行任意可执行程序                                   |
| **生成报告** | `mvn site`                                | 生成项目站点                                         |
|              | `mvn site:deploy`                         | 部署生成的网站                                       |
| **发布管理** | `mvn release:prepare`                     | 准备一个发布版本，包含打标签和增加版本号             |
|              | `mvn release:perform`                     | 执行发布，包含将项目打包并部署到远程仓库             |
| **代码质量** | `mvn checkstyle:check`                    | 运行Checkstyle插件来检查代码风格                     |
|              | `mvn pmd:check`                           | 运行PMD插件来检查代码问题                            |
|              | `mvn findbugs:check`                      | 运行FindBugs插件来检查代码中的潜在错误               |
| **文档生成** | `mvn javadoc:javadoc`                     | 生成项目的Javadoc文档                                |
| **依赖信息** | `mvn dependency:resolve`                  | 解析并显示项目的依赖信息                             |
|              | `mvn dependency:tree`                     | 以树形结构显示项目的依赖关系                         |
| **版本信息** | `mvn versions:display-dependency-updates` | 显示项目依赖的最新版本信息                           |
|              | `mvn versions:display-plugin-updates`     | 显示项目插件的最新版本信息                           |
| **其他**     | `mvn help:help`                           | 显示帮助信息                                         |
|              | `mvn versions:set -DnewVersion=version`   | 设置项目的新版本号                                   |

这张表格涵盖了Maven在项目管理、依赖管理、插件管理、运行项目、生成报告、发布管理、代码质量、文档生成、依赖信息、版本信息等方面的常用命令，帮助开发者更全面地了解和使用Maven的命令行工具。