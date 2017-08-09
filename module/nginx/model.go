package nginx

type Model struct{
  Filename string `json:"filename" form:"filename" query:"filename"`
  Content string `json:"content" form:"content" query:"content"`
}

