# TraefikIpFirewall
Traefikのfoward authで利用できる、IPベースのアクセス許可サービス


# 動作確認


## rule.yaml のサンプル

```yaml:rule.yaml
black_list:
- 192.168.1.1/32
```

## X-Forwarded-For がrule.yamlにマッチする値の場合

```bash
$ curl -I -H X-Forwarded-For:192.168.1.1 localhost:8081
HTTP/1.1 401 Unauthorized
```

## X-Forwarded-For がrule.yamlにマッチしない場合値の場合

```bash
curl -I -H X-Forwarded-For:192.168.1.2 localhost:8081
HTTP/1.1 202 Accepted
```

## X-Forwarded-For にIPアドレス以外が設定されていた場合

```bash
$ curl -I -H X-Forwarded-For:test localhost:8081 
HTTP/1.1 401 Unauthorized
```
