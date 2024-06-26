<template>
  <div class="login-form">
    <form @submit.prevent="handleLogin">
      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'LoginForm',
  setup() {
    const email = ref<string>('')
    const password = ref<string>('')

    const handleLogin = async () => {
      try {
        const { data, error } = await useFetch('http://localhost:8080/login', {
          method: 'POST',
          body: { email: email.value, password: password.value },
          headers: { 'Content-Type': 'application/json' },
        })
        if (error.value) {
          throw new Error(error.value.message)
        }
        console.log('Login successful:', data.value)
      } catch (error) {
        console.error('Login failed:', error)
      }
    }

    return { email, password, handleLogin }
  },
})
</script>

<style scoped>
.login-form {
  max-width: 300px;
  margin: auto;
  padding: 1em;
  border: 1px solid #ccc;
  border-radius: 5px;
}
div {
  margin-bottom: 1em;
}
label {
  margin-bottom: 0.5em;
  color: #333333;
  display: block;
}
input {
  border: 1px solid #cccccc;
  padding: 0.5em;
  font-size: 1em;
  border-radius: 3px;
  width: 100%;
}
button {
  padding: 0.7em;
  color: #fff;
  background-color: #007bff;
  border: none;
  border-radius: 3px;
  cursor: pointer;
}
button:hover {
  background-color: #0056b3;
}
</style>
