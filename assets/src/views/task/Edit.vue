<template>
  <page-header-wrapper :title="false">
    <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
      <a-form @submit="handleSubmit" :form="form">
        <a-form-item
          label="任务名称"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"
          help="任务名称"
        >
          <a-input
            v-decorator="[
              'name',
              {rules: [{ required: true, message: '请输入任务名称' }], initialValue: task.name},
            ]"
            name="name"
            placeholder="" />
        </a-form-item>
        <a-form-item
          v-model="task.spec"
          label="cron表达式"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"
          help="秒 分 时 天 周 年"
        >
          <a-input
            name="spec"
            placeholder="* * * * * *"
            v-decorator="[
              'spec',
              {rules: [{ required: true, message: '请输入cron表达式' }], initialValue: task.spec}
            ]" />
        </a-form-item>
        <a-form-item
          label="执行命令"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"
          help="任务执行的命令"
        >
          <a-textarea
            rows="6"
            placeholder="请输入执行命令"
            name="command"
            v-decorator="[
              'command',
              {rules: [{ required: true, message: '请输入执行命令' }], initialValue: task.command}
            ]" />
        </a-form-item>
        <a-form-item
          label="超时时间"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"
          help="最少不能小于 1 秒"
        >
          <a-input-number
            :min="1"
            name="timeout"
            v-decorator="[
              'timeout',
              {rules: [{ required: false }], initialValue: task.timeout}
            ]"
          />
          <span> 秒</span>
        </a-form-item>
        <a-form-item
          label="重试次数"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"
          help="0 表示失败不进行重试"
        >
          <a-input-number
            :min="0"
            name="retryTimes"
            v-decorator="[
              'retryTimes',
              {rules: [{ required: false }], initialValue: task.retry_times}
            ]"
          />
          <span> 次</span>
        </a-form-item>

        <a-form-item
          label="重试间隔"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17}}"
          :required="false"
          help="重试次数不为 0 时必填"
        >
          <a-input-number
            :min="0"
            name="retryInterval"
            v-decorator="[
              'retryInterval',
              {rules: [{ required: false }], initialValue: task.retry_interval}
            ]"
          />
          <span> 秒</span>
        </a-form-item>

        <a-form-item
          label="状态"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"
          :required="false"
        >
          <a-radio-group v-decorator="['status', { initialValue: task.status }]">
            <a-radio :value="10">运行</a-radio>
            <a-radio :value="20">关闭</a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item
          :wrapperCol="{ span: 24 }"
          style="text-align: center"
        >
          <a-button htmlType="submit" type="primary">保存</a-button>
          <a-button style="margin-left: 8px" @click="goList">返回</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import { editTask, getTaskDetail } from '@/api/task'

export default {
  name: 'TaskEdit',
  data () {
    return {
      form: this.$form.createForm(this),
      task: {}
    }
  },
  mounted () {
    this.task.id = this.$route.params.id
    this.getDetail()
  },
  methods: {
    handleSubmit (e) {
      e.preventDefault()
      this.form.validateFields((err, values) => {
        values.id = this.task.id
        if (!err) {
          editTask(values).then((res) => {
            this.goList()
          })
        }
      })
    },

    getDetail () {
      getTaskDetail(this.task.id).then(res => {
        this.task = res.result.data
      })
    },

    goList () {
      this.$router.push({ name: 'TaskList' })
    }
  }
}
</script>
