<template>
  <main class="p-6 bg-gray-100 min-h-screen">
    <div class="max-w-7x1 mx-auto bg-white p-6 rounded-lg shadow-lg">
      <div class="flex justify-between items-center mb-6">
        <h4 class="text-xl font-semibold text-gray-700">
          Transactions
        </h4>
        <RouterLink
            to="/transaction/create"
            class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg shadow-md"
        >
          Add transaction
        </RouterLink>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full bg-white-border border-gray-300 rounded-lg">
          <thead>
          <tr class="bg-gray-200 text-gray-700 uppercase text-sm leading-normal">
            <th class="py-3 px-6 text-left">ID</th>
            <th class="py-3 px-6 text-left">Type</th>
            <th class="py-3 px-6 text-left">Ticker</th>
            <th class="py-3 px-6 text-left">Volume</th>
            <th class="py-3 px-6 text-left">Price</th>
            <th class="py-3 px-6 text-left">Date</th>
            <th class="py-3 px-6 text-left">Actions</th>
          </tr>
          </thead>
          <tbody v-if="transactions.length > 0">
          <tr
              v-for="(transaction, index) in transactions"
              :key="index"
              class="border-b border-gray-300 hover:bg-gray-50"
          >
            <td class="py-3 px-6 text-black">{{ transaction.id }}</td>
            <td class="py-3 px-6 text-black">{{ transaction.type }}</td>
            <td class="py-3 px-6 text-black">{{ transaction.ticker }}</td>
            <td class="py-3 px-6 text-black">{{ transaction.volume }}</td>
            <td class="py-3 px-6 text-black">{{ transaction.price }}</td>
            <td class="py-3 px-6 text-black">{{ transaction.date }}</td>
            <td class="py-3 px-6 flex space-x-2">
              <RouterLink
                  :to="{ path: '/transaction/' + transaction.id + '/edit'}"
                  class="bg-green-500 hover-bg-green-600 text-white font-semibold py-2 px-4 rounded-lg shadow-lg"
              >
                Edit
              </RouterLink>
              <button
                  @click="deleteTransaction(transaction.id)"
                  type="button"
                  class="bg-red-500 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-lg shadow-lg"
              >
                Delete
              </button>
            </td>
          </tr>
          </tbody>
          <tbody v-else>
          <tr>
            <td class="py-3 px-6 text-center text-gray=500" colspan="7">
              Loading...
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import axios from "axios";

interface Transaction {
  id: number;
  type: string;
  ticker: string;
  volume: number;
  price: number;
  date: string;
}

export default defineComponent({
  name: 'dashboard',
  data() {
    return {
      transactions: [] as Transaction[],
    }
  },
  mounted() {
    this.getTransactions();
  },
  methods: {
    getTransactions() {
      axios.get('http://localhost:80/api/txn').then(res => {
        this.transactions = res.data.transactions
        console.log("this.transactions")
        console.log(this.transactions)
      }).catch(error => {
        console.error("Error fetching transactions:", error);
      })
    },
    deleteTransaction(studentId: number) {
      if (confirm('Are you sure you want to delete transaction?')) {
        axios.delete(`http://localhost:80/api/txn/delete`, {
          data: {
            id: Number(studentId)
          }
        }).then(res => {
          alert(res.data.message)
          this.getTransactions()
        }).catch(error => {
          if (error.response) {
            alert(error.response.data.error)
          }
        })
      }
    }
  }
})
</script>