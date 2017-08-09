package nginx 

type Controller struct {}

// GetAllConfigFilename is a method to
// get the list of nginx configuration file according to
// path configured by user in config.toml
func (c *Controller) GetAllConfigFilename() (interface{}, error) {
  logic := new(Logic)
  filename := logic.FindAll()
  return filename, nil
}

// ReadConfigFile is a method to
// get a config file by filename and read its content
func (c *Controller) ReadConfigFile(filename string) (string, error) {
  logic := new(Logic)
  file := logic.FindOne(filename)
  return file, nil
}

// CreateConfigFile is a method to 
// create a new nginx config file
func (c *Controller) CreateConfigFile(filename string) (interface{}, error) {
  logic := new(Logic)
  err := logic.Create(filename)
  return map[string]string{"file": filename}, err
}

// UpdateConfigFile is a method to
// update an existing nginx config file by filename
func (c *Controller) UpdateConfigFile(filename string, content string) (interface{}, error) {
  logic := new(Logic)
  err := logic.Update(filename, content)
  return map[string]string{"file": filename, "content": content}, err
}