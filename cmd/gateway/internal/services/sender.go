package services

func PackageSender(provider *Provider) {
	for {
		pack := <-provider.PackChan

		provider.Log.Sugar().Infof("package to be sent: %+v", pack)
	}
}
