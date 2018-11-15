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
                <q-input v-if="modalType == 'login'" v-model="user.login" float-label="Логин или email*" clearable/>
                <q-input v-if="modalType == 'signup'" v-model="user.nickname" float-label="Логин*" clearable/>
                <q-input v-if="modalType == 'signup'" v-model="user.email" float-label="Email*" clearable/>
                <q-input v-model="user.password" float-label="Пароль*" clearable/>
                <q-input v-if="modalType == 'signup'" v-model="user.name" float-label="Имя" clearable/>
                <q-input v-if="modalType == 'signup'" v-model="user.lastname" float-label="Фамилия" clearable/>
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
  import * as utils from './../utils'
  //todo: validate user inputs
  export default {
    name: "LoginForm",
    data () {
    return {
        opened: false,
        modalType: "",
        user: {
          login: "",
          email: "",
          phone: "",
          nickname: "",
          name: "",
          lastname: ""
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
      this.ClearUser();
    },
    methods: {
    Login: function () {
      this.axios.post('session',{
            login: this.user.login,
            password: this.user.password
        }, {
            jar:true,
            withCredentials: true
        })
        .then(response => {
          this.opened = false;
          this.ClearUser();
        })
        .catch(e => {
          console.log(e.response);
          if (e.response.status === 401) {
            utils.notify401();
          } else {
            utils.notify500();
          }
            this.opened = true;
          })
    },
    SignUp: function () {
      console.log("SIGNUP: " + this.user)
    },
    ClearUser: function () {
      for (let key in this.user) {
        if (this.user.hasOwnProperty(key)) {
          this.user[key] = "";
        }
      }
    }
    }
  }
</script>

<style scoped>

</style>