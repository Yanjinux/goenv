/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-19 00:38:24
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-06-09 23:36:39
 * @FilePath: \code\village\common\xerr\errCode.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package xerr

//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005

//用户模块
const MSG_SEND_FREQUENCE uint32 = 200001 // 短息发送频繁
