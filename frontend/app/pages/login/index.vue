<!-- pages/login.vue -->
<template>
  <div class="login-page">
    <h1>Login</h1>
    <LoginForm />
  </div>
</template>

<script lang="ts">
import LoginForm from '../../components/LoginForm.vue'

export default {
  components: {
    LoginForm: LoginForm,
  },
}
</script>

<style scoped>
.login-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}
</style>
