# oscli
## elasticsearch command-line tool

### simple app in golang to give easy access to elasticsearch api calls from command line based on your iam role in AWS


### Requirements:
  - must have aws cli installed and configured with permissions to elasticsearch in AWS
  - elasticsearch endpoint URI exposed and known
  - go build requires access to the internet to download packages

Deployment:
clone this repository
```console
git clone {.git url}
```

```console 
go build
```


```console
./oscli
usage:

admin
	 ls-role-mapping | ls-tenants | ls-security
configure
	 { allows for aws profile and elastic host url }
role-mapping
	 new -name=role_name -user=elasticsearch_user -backend-role=kibana_user -host={nil or host match}
indice
	 delete -name=indexName
template
	 view -name=templateName
role
	 new -name=test_role -clusterperms=indices_monitor,cluster_ops -index="movies-*" -indexperm=read -tenant=sales,marketing -tenantperms=kibana_all_read
cat
	 cluster-info | stats | nodes | templates | indices | shards | health | disk | recovery | master | count | field_data | alias
help: { list this message :-) }
```
example
```console
some-user@ww2r32342de: ./oscli configure
AWS Cli Profile [default]:
ElasticSearch Host [https://my-elasticsearch-host-url.us-east-1.es.amazonaws.com]:
config file written..
some-user@ww2r32342de:
```

```console
some-user@ww2r32342de: ./oscli configure
AWS Cli Profile [default]: us-east-1-dev
ElasticSearch Host [https://my-elasticsearch-host-url.us-east-1.es.amazonaws.com]: http://my-dev-cluster.us-east-1.es.amazonaws.com
config file written..
some-user@ww2r32342de:
```

```console
./oscli admin ls-role-mapping
```



