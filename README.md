# cloudeye-grafana
[![LICENSE](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://github.com/huaweicloud/cloudeye-grafana/blob/master/LICENSE)

Сloudeye-grafana is a datasource plugin to adapt to Grafana.

# Quick start

## 1. Installation
> Preparation before installation:
> a. Grafana version >=7.4.0, [grafana](https://grafana.com/grafana/download)  
> b. Download cloudeye-grafana.tar.gz from[release](https://github.com/huaweicloud/cloudeye-grafana/releases)cloudeye-grafana.tar.gz

### 1.1 从release安装
a. Place the downloaded plugin package into the grafana plugin directory (see the plugins configuration path in conf/defaults.ini),
decompress cloudeye-grafana-{version}.tar.gz, and pay attention to maintaining directory permissions and grafana running permissions.
  
b. Modify conf/defaults.ini to allow unsigned plugins to run 
> allow_loading_unsigned_plugins = cloudru-cloudeye-grafana  
   
c. Restart grafana

## 2. Configure cloudeye-grafana data source
> Preparation before configuration: 
> a. [Obtain AK/SK](https://support.hc.sbercloud.ru/devg/apisign/api-sign-provide-aksk.html)  
> b. [Obtain project_id](https://support.hc.sbercloud.ru/devg/apisign/api-sign-provide-proid.html)  
> c. [Obtain CES Endpoint and RegionID](https://support.hc.sbercloud.ru/endpoint/index.html)

### 2.1 Configure data source
a. Enter grafana's data source configuration page (Data Sources), click Add data source to enter the configuration form page,
fill in the data source name cloudeye-grafana, and select cloudeye-grafana in the data source list.

b. Currently, supports only Specific Region Mode

> Specific Region Mode：configure CES Endpoint、Region ID、Project ID、IAM Access Key、IAM Secret Key

c. (Optional) If you need to enable reading indicator metadata through the configuration file, you need to click
the Get Metric Meta From Conf button to enable it, and configure the indicator metadata [list as follows](#metrics).

d. Click the Save & test button. If Data source is working is displayed, it means that the data source is configured successfully and you can start accessing Cloudru monitoring data in grafana.


<a name="metrics"></a>
## 3. (Optional) Configure indicator metadata list
In order to improve the query experience, for tenants whose resource list changes in real time and has a large amount of resources,
the resource list can be configured in the dist/metric.yaml file in advance. The region/service/resource/metric list shall be subject to the configuration file.  
> a. [云监控支持的服务指标列表](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)     
> b. [华为云支持region列表](https://developer.huaweicloud.com/endpoint)  
> c. 按metric.yaml样例配置完成后，重启grafana
    
## 4. 导入dashboard模板
为简便租户配置，本插件提供了ECS、ELB、RDS服务的Dashboard预设模板，见： cloudeye-grafana/src/templates目录

## 5. 创建自定义Dashboard
a. 鼠标移动至页面左侧菜单"+"图标，选择Dashboard，点击即可创建

b. 创建好之后请点击右上角齿轮图标，选择左侧"Variables"菜单项，点击"Add variable"按钮添加filter和period模板变量。变量配置如下:
> filter变量：
```
{
    Name: filter,
    Type: Query,
    Label: Filter,
    Data source: cloudeye-grafana,
    Query: listFilterOptions()
}
```

> period变量：
```
{
    Name: period,
    Type: Query,
    Label: Period,
    Data source: cloudeye-grafana,
    Query: listPeriodOptions()
}
```

c. 配置好自定义模板变量后回到Dashboard页面，点击"Add an empty panel"按钮添加指标监控图表

d. 点击右上角保存按钮，完成自定义Dashboard创建