### How to import golang local package in module mode

```
https://medium.com/faun/golang-package-management-using-go-modules-d3c929900114
```

#### In root, open terminal and run
```
go mod init jacky.com/dddtest
```

#### Then, no need to create any mod file in sub folder, just import it by this absolute path
##### root module name / folder 
```
import (
	"jacky.com/dddtest/patient"
	"jacky.com/dddtest/twoLayerFolder/subtwolayerfolder"
)
```

#### Remember, the folder name and package name must be the same !!!
