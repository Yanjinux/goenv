/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-19 00:38:51
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-19 00:39:00
 * @FilePath: \code\village\common\xerr\errors.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package xerr

import (
	"fmt"
)

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
