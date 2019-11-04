package errs

import "backend/serializer"

const (
	// 全局
	SUCCESS        = 0
	ERROR          = 500
	INVALID_PARAMS = 400

	// 鉴权相关
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERR_EMAIL                      = 20005
	ERR_USER_NAME                  = 20006
	ERR_GEN_PASSWD                 = 20007
	ERR_REGISTER                   = 20008

	ERR_WRONG_USER_PASSWD = 20009

	// 文章相关
	ERROR_EXIST_TAG       = 10001
	ERROR_EXIST_TAG_FAIL  = 10002
	ERROR_NOT_EXIST_TAG   = 10003
	ERROR_GET_TAGS_FAIL   = 10004
	ERROR_COUNT_TAG_FAIL  = 10005
	ERROR_ADD_TAG_FAIL    = 10006
	ERROR_EDIT_TAG_FAIL   = 10007
	ERROR_DELETE_TAG_FAIL = 10008
	ERROR_EXPORT_TAG_FAIL = 10009
	ERROR_IMPORT_TAG_FAIL = 10010

	ERROR_NOT_EXIST_ARTICLE        = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012
	ERROR_ADD_ARTICLE_FAIL         = 10013
	ERROR_DELETE_ARTICLE_FAIL      = 10014
	ERROR_EDIT_ARTICLE_FAIL        = 10015
	ERROR_COUNT_ARTICLE_FAIL       = 10016
	ERROR_GET_ARTICLES_FAIL        = 10017
	ERROR_GET_ARTICLE_FAIL         = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:           "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_GET_TAGS_FAIL:            "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:           "统计标签失败",
	ERROR_ADD_TAG_FAIL:             "新增标签失败",
	ERROR_EDIT_TAG_FAIL:            "修改标签失败",
	ERROR_DELETE_TAG_FAIL:          "删除标签失败",
	ERROR_EXPORT_TAG_FAIL:          "导出标签失败",
	ERROR_IMPORT_TAG_FAIL:          "导入标签失败",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_ADD_ARTICLE_FAIL:         "新增文章失败",
	ERROR_DELETE_ARTICLE_FAIL:      "删除文章失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL: "检查文章是否存在失败",
	ERROR_EDIT_ARTICLE_FAIL:        "修改文章失败",
	ERROR_COUNT_ARTICLE_FAIL:       "统计文章失败",
	ERROR_GET_ARTICLES_FAIL:        "获取多个文章失败",
	ERROR_GET_ARTICLE_FAIL:         "获取单个文章失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL:  "生成文章海报失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERR_EMAIL:                      "邮箱已被注册",
	ERR_USER_NAME:                  "用户名已被注册",
	ERR_GEN_PASSWD:                 "密码生成失败",
	ERR_REGISTER:                   "注册失败",
	ERR_WRONG_USER_PASSWD: "用户名或密码错误",

}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// 序列化错误信息, 出错时返回错误码和错误信息，其他不返回
func BuildErrorResponse(status int) *serializer.Response {
	return &serializer.Response{Status: status, Error: GetMsg(status)}
}