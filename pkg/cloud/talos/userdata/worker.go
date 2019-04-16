package userdata

const workerTempl = `---
version: ""
security: null
services:
  init:
    cni: flannel
  kubeadm:
    configuration: |
      apiVersion: kubeadm.k8s.io/v1beta1
      kind: JoinConfiguration
      discovery:
        bootstrapToken:
          token: '{{ .KubeadmTokens.BootstrapToken }}'
          unsafeSkipCAVerification: true
          apiServerEndpoint: {{ index .MasterIPs 0 }}:443
      nodeRegistration:
        taints: []
        kubeletExtraArgs:
          node-labels: ""
          feature-gates: ExperimentalCriticalPodAnnotation=true
      token: '{{ .KubeadmTokens.BootstrapToken }}'
  trustd:
    username: '{{ .TrustdInfo.Username }}'
    password: '{{ .TrustdInfo.Password }}'
    endpoints: [ {{ range $i,$ip := .MasterIPs }}{{if $i}},{{end}}"{{$ip}}"{{end}} ]
`