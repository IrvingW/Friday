<template>
<el-container>
  <el-container>
    <el-header>
      <!-- Header content -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
          <el-breadcrumb-item :to="{ path: '/friday/home/resource_view' }">资源视图</el-breadcrumb-item>
          <el-breadcrumb-item> </el-breadcrumb-item>
        </el-breadcrumb>
    </el-header>
    <el-main>
      <!-- Main content -->
      <el-card shadow="hover" :body-style="{ padding: '20px' }">
        <el-row class="select-row">
          <el-col :span="2"> <p>集群名称：</p> </el-col>
          <el-col :span="12">
              <el-select v-model="select_value" filterable placeholder="请选择集群" class="select">
                <el-option
                  v-for="item in select_options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <el-tabs active-name="resource">
              <el-tab-pane label="资源视图" name="resource">
                <el-row>
                  <el-col :span="23">
                      <el-row :gutter="30">
                        <el-col :span="6">
                          <ve-pie :data="cpu_usage"></ve-pie>
                        </el-col>
                        <el-col :span="6">
                          <ve-pie :data="memory_usage"></ve-pie>
                        </el-col>
                        <el-col :span="6">
                          <ve-pie :data="disk_usage"></ve-pie>
                        </el-col>
                      </el-row>
                      <el-row :gutter="10" type="flex" justify="center">
                        <el-col :span="6"><p class="title">CPU使用率</p></el-col>
                        <el-col :span="6"><p class="title">内存使用率</p></el-col>
                        <el-col :span="6"><p class="title">磁盘使用率</p></el-col>
                      </el-row>
                      </el-col>
                    </el-row>
                </el-tab-pane>
              <el-tab-pane label="节点列表" name="list">
                <el-table :data="node_list" stripe>
                  <el-table-column v-for="col in columns"
                    :prop="col.id"
                    :key="col.id"
                    :label="col.label"
                    :width="col.width">
                  </el-table-column>
                  <el-table-column label="状态" width="120">
                      <template slot-scope="scope">
                        <el-tag :type="scope.row.status === 'error' ? 'error' : 'success'"
                          disable-transitions>{{scope.row.status === 'error' ? '异常' : '正常'}}</el-tag>
                      </template>
                  </el-table-column>
                </el-table>
              </el-tab-pane>
            </el-tabs>
          </el-col>
        </el-row>
      </el-card>
    </el-main>
  </el-container>
</el-container>

</template>
<script>
  export default {
	name:'node_view',
	  data(){
		return{
      cpu_usage: {
        columns: ['name', 'percent'],
        rows: [
          {'name': '未使用(%)', 'percent': 87.2},
          {'name': '已使用(%)', 'percent': 12.8}
        ]
      },
      memory_usage: {
        columns: ['name', 'count'],
        rows: [
          {'name': '未使用(G)', 'count': 98.3},
          {'name': '已使用(G)', 'count': 1.7}
        ]
      },
      disk_usage: {
        columns: ['name', 'count'],
        rows: [
          {'name': '未使用(G)', 'count': 49},
          {'name': '已使用(G)', 'count': 51}
        ]
      },
      node_list:[
        {ip: "192.168.1.116", role: "master", cpu_cnt: 4, cpu_usage: 24, memory_usage: 2, disk_usage: 27, status: "normal"},
        {ip: "192.168.1.108", role: "slave", cpu_cnt: 8, cpu_usage: 7, memory_usage: 1, disk_usage: 43, status: "normal"},
        {ip: "192.168.1.194", role: "slave", cpu_cnt: 12, cpu_usage: 13, memory_usage: 1, disk_usage: 83, status: "normal"}
      ],
      columns: [
        {id: "ip", label: "IP", "width": 150},
        {id: "role", label: "角色", "width": 150},
        {id: "cpu_cnt", label: "CPU核数", "width": 150},
        {id: "cpu_usage", label: "CPU使用率(%)", "width": 330},
        {id: "memory_usage", label: "内存使用率(%)", "width": 120},
        {id: "disk_usage", label: "磁盘使用率(%)", "width": 120},
      ],
      select_options: [{
        value: '选项1',
        label: 'dclab'
      }],
      select_value: ''
		}
    }
  }
</script>
<style type="text/css" scoped>
.select {
  margin-top: 10px;
}
.select-row {
  margin-bottom: 30px;
}
</style>