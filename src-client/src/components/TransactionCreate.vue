<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  name: 'transactionCreate',
  date() {
    return {
      err: null,
      model: {
        transaction: {
          type: '',
          ticker: '',
          volume: '',
          price: '',
          date: '',
        }
      }
    }
  }
})


</script>