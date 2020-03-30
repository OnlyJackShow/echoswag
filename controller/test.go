package controller

import "github.com/labstack/echo/v4"

// @Summary 品牌列表
// @Tags 商家品牌
// @Accept plain
// @Produce json
// @Param id query string false "品牌ID"
// @Param name query string false "品牌名称"
// @Param initial query string false "品牌首字母"
// @Param category_id query string false "品牌分类ID"
// @Param page_index query int true "分页索引"
// @Param page_size query int true "分页大小"
// @Success 200 {object} dto.BaseResponse
// @Failure 400 {object} dto.BaseResponse
// @Router /boss/product/brand/list [get]
func QueryBrand(e echo.Context)error {
	return e.String(200, "ok")
}


