package schema

import "github.com/dashenwo/go-backend/v2/console/category/proto"

// 添加一条数据请求参数
type CategoryCreateParam proto.AddRequest

// 添加一条数据请求参数
type CategoryCreateResult proto.AddResponse

// 查询一条参数
type CategoryQueryOneParam proto.QueryOneRequest

// 查询一条的返回
type Category proto.CategorySchema

// 查询所有对的参数
type CategoryQueryParam proto.QueryRequest

// 查询所有的参数返回
type CategoryQueryResult proto.QueryResponse

// 修改数据的请求参数
type CategoryEditParam proto.EditRequest

// 修改数据的返回
type CategoryEditResult proto.EditResponse

// 删除一条数据请求参数
type CategoryDeleteParam proto.DeleteRequest

// 删除一条数据请求参数
type CategoryDeleteResult proto.DeleteResponse
