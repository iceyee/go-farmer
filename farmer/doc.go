// 一些常用的功能, 包括断言, 异常, 日志, Http, 信号量, StringBuilder.
//
/*
日志系统的使用:
  1. 调用OpenLog()开启日志系统.
  2. 调用Info(), Warn(), Error()来写日志, 日志会输出到终端以及文件.
  3. 需要关闭日志系统时则调用CloseLog().
  // 最后日志会写在"/opt/farmer-log/项目名/项目名-WARN.log"和"/opt/farmer-log/项目名/项目名-ERROR.log".
  // Info(), Warn(), Error()等接口接收的参数可以是string, error, fmt.Stringer, 其它类型则会输出[Unkown].

查看本机ip:
  http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json
*/
package farmer
