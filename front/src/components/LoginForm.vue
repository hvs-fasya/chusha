<template>
    <q-modal v-model="opened">
        <q-modal-layout>
            <q-toolbar slot="header" inverted>
                <q-btn
                        flat
                        round
                        dense
                        v-close-overlay
                        icon="keyboard_arrow_left"
                />
                <q-toolbar-title>
                    <span v-if="modalType === 'login'">Вход</span>
                    <span v-if="modalType === 'signup'">Регистрация</span>
                </q-toolbar-title>
            </q-toolbar>

            <div class="layout-padding">
                <q-input v-if="modalType == 'login'" v-model="user.nickname" float-label="Логин или email" clearable/>
                <q-input v-if="modalType == 'signup'" v-model="user.nickname" float-label="Логин" clearable/>
                <q-input v-if="modalType == 'signup'" v-model="user.email" float-label="Email" clearable/>
                <q-input v-model="user.password" float-label="Пароль" clearable/>
                <div class="q-mt-xl">
                    <q-btn
                            rounded
                            outline
                            class="q-mr-md"
                            color="primary"
                            v-close-overlay
                            label="Отмена"
                    />
                    <q-btn
                            rounded
                            class="q-ml-md"
                            color="primary"
                            v-close-overlay
                            label="Отправить"
                            @click="modalType == 'login' ? Login() : SignUp()"
                    />
                </div>
            </div>
        </q-modal-layout>
    </q-modal>
</template>

<script>
  import { EventBus } from './../event-bus.js'
  //todo: validate user inputs
  export default {
    name: "LoginForm",
    data () {
    return {
        opened: false,
        modalType: "",
        user: {
          nickname: "",
          password: "",
          email: ""
        }
      }
    },
    mounted(){
      EventBus.$on('user-form-open', (type) => {
        this.modalType = type;
        this.opened = true;
      })
    },
    beforeDestroy(){
      EventBus.$off('user-form-open');
      this.opened = false;
      this.modalType = "";
      this.user = {
        nickname: "",
        password: "",
        email: ""
      }
    },
    methods: {
    Login: function () {
      console.log("LOGIN: " + this.user);
      this.axios.post('login',{
            username: this.user.nickname,
            password: this.user.password
        },
        {
          withCredentials: true
        })
        .then(response => {
            console.log(response.data)
        })
        .catch(e => {
            console.log("ERROR: " + e)
          })
    },
    SignUp: function () {
      console.log("SIGNUP: " + this.user)
    }
    }
  }
</script>

<style scoped>

</style>