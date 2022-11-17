<template>
  <page-header-wrapper :title="false">
    <a-card :body-style="{padding: '10px 10px'}" :bordered="false">
      <a-table
        :columns="columns"
        :data-source="items"
        rowKey="id"
        :pagination="pagination"
        @change="handlePageChange"
        :loading="loading"
      >
        <span slot="status" slot-scope="text, record">
          <a-tag v-if="record.status === 2" color="red">
            失败
          </a-tag>
          <a-tag v-if="record.status === 0" color="orange">
            执行中
          </a-tag>
          <a-tag v-if="record.status === 1" color="green">
            完成
          </a-tag>
        </span>
        <span slot="execTime" slot-scope="text, record">
          执行时长: {{ record.total_time > 0 ? scope.row.total_time : 1 }}秒<br>
          开始时间: {{ record.start_time | formatTime }}<br>
          <span v-if="record.status !== 1">结束时间: {{ record.end_time | formatTime }}</span>
        </span>
        <span slot="action" slot-scope="text, record">
          <template>
            <a-button @click="showTaskResult(record)" type="primary">查看结果</a-button>
          </template>
        </span>
      </a-table>
      <a-modal
        title="任务执行结果"
        :visible="dialogVisible"
        :maskClosable="true"
        @cancel="handleCancel"
      >
        <template slot="footer">
          <a-button key="back" type="danger" @click="handleCancel">
            关闭
          </a-button>
        </template>
        <div>
          <pre>{{ this.command }}</pre>
        </div>
        <div>
          <pre>{{ this.result }}</pre>
        </div>
      </a-modal>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { taskLogList } from '@/api/task'

export default {
  name: 'TaskLogList',
  components: {
  },
  data () {
    return {
      // 查询参数
      queryParam: {},
      // 表头
      columns: [
        {
          title: 'ID',
          dataIndex: 'id',
          align: 'center',
          key: 'id'
        },
        {
          title: '任务ID',
          dataIndex: 'task_id',
          align: 'center'
        },
        {
          title: '任务名称',
          dataIndex: 'name',
          align: 'center',
          width: '200px'
        },
        {
          title: '执行状态',
          dataIndex: 'status',
          align: 'center',
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '执行时长',
          dataIndex: 'execTime',
          align: 'center',
          scopedSlots: { customRender: 'execTime' }
        },
        {
          title: '执行结果',
          dataIndex: 'action',
          align: 'center',
          width: '150px',
          scopedSlots: { customRender: 'action' }
        }
      ],
      items: [],
      selectedRowKeys: [],
      selectedRows: [],
      loading: false,
      pagination: {
        current: 1,
        total: 0,
        pageSize: 0
      },
      taskId: 0,
      dialogVisible: false,
      command: '',
      result: ''
    }
  },
  created () {
    this.taskId = this.$route.params.id
    this.taskLogList()
  },
  methods: {
    handleEdit (record) {
      this.$router.push({ name: 'TaskEdit', params: { id: record.id } })
    },

    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },

    resetSearchForm () {
      this.queryParam = {
        date: moment(new Date())
      }
    },

    taskLogList () {
      this.loading = true
      taskLogList(this.taskId, this.queryParam).then(res => {
        this.items = res.result.items
        this.pagination.current = res.result.pager.page
        this.pagination.total = res.result.pager.total_rows
        this.pagination.pageSize = res.result.pager.page_size
        this.loading = false
      })
    },

    handlePageChange (pagination) {
      this.pagination.current = pagination.current
      this.queryParam.page = pagination.current
      this.taskLogList()
    },

    goTaskLog () {
      this.$router.push({ name: 'TaskLogList' })
    },

    handleCancel () {
      this.dialogVisible = false
    },

    handleOk () {
      this.dialogVisible = false
    },

    showTaskResult (task) {
      this.dialogVisible = true
      this.command = task.command
      this.result = task.result
    }
  }
}
</script>
