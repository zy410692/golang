package zstackapi

import "github.com/gzxgogh/zstack-sdk-go-v4/zstack"

func main() {
	// zstack/ZStackVm.go
	// 免费版本的zstack没有看到生成accesskeyid的地方,在zstack-cli里面还有
	// admin >>>CreateAccessKey 36c27e8ff05c4780bf6d2fa65700f22e  36c27e8ff05c4780bf6d2fa65700f22e
	//Invalid parameter[36c27e8ff05c4780bf6d2fa65700f22e], the parameter must be split by "="
	//admin >>>CreateAccessKey accountUuid=36c27e8ff05c4780bf6d2fa65700f22e  36c27e8ff05c4780bf6d2fa65700f22e
	//Invalid parameter[36c27e8ff05c4780bf6d2fa65700f22e], the parameter must be split by "="
	//admin >>>CreateAccessKey accountUuid=36c27e8ff05c4780bf6d2fa65700f22e userUuid=36c27e8ff05c4780bf6d2fa65700f22e
	//API call[org.zstack.header.message.APIEvent] failed because [code: LICENSE.1001, description: The license is not permitted for the operation, details: API[class org.zstack.accessKey.APICreateAccessKeyMsg] is not allowed for the community-source license, please apply an enterprise license]

	//原来还需要license支持  --

	zstack.CreateVmInstances()

}
