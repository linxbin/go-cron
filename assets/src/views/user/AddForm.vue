<template>
  <a-modal
    :visible="visible"
    title="新增用户"
    okText="提交"
    @cancel="() => { $emit('cancel') }"
    @ok="() => { $emit('create') }"
  >
    <a-form layout="vertical" :form="form">
      <a-form-item label="用户名">
        <a-input
          v-decorator="[
            'username',
            {rules: [{ required: true, message: '请输入用户名' }]}
          ]"
        />
      </a-form-item>
      <a-form-item label="密码">
        <a-input
          type="password"
          v-decorator="[
            'password',
            {rules: [{ required: true, message: '请输入密码' }]}
          ]"
        />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script>
export default {
  name: 'AddForm',
  props: ['visible'],
  beforeCreate () {
    this.form = this.$form.createForm(this, { name: 'form_in_modal' })
  }
}
</script>
