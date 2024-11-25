<template>
    <el-container class="container">
        <el-form :model="form" class="exchange-form">
            <el-form-item label="从货币" label-width="80px">
            <el-select v-model="form.tocurrency" placeholder="选择货币">
                <el-option v-for="currency in currencies" :key="currency" :lable="currency" :value="currency"/>
            </el-select>
            </el-form-item>
            <el-form-item label="到货币" label-width="80px">
            <el-select v-model="form.fromcurrency" placeholder="选择货币">
                <el-option v-for="currency in currencies" :key="currency" :lable="currency" :value="currency"/>
            </el-select>
            </el-form-item>
            <el-form-item label="金额" label-width="80px">
            <el-input v-model="form.amount" type="number" placeholder="输入金额"/>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="exchange">兑换</el-button>
            </el-form-item>            
        </el-form>
        <div class="result">
            <p>兑换结果:{{result}}</p>
        </div>
    </el-container>
</template>

<script setup lang="ts">
import {ref, onMounted} from 'vue';
import axios from '../axios';

interface ExchangeRate{
    fromcurrency: string;
    tocurrency: string;
    rate: number;
}

const form = ref({
    fromcurrency: '',
    tocurrency: '',
    amount: 0
});

const result = ref<number | null>(null);
const currencies = ref<string[]>([]);
const rates = ref<ExchangeRate[]>([]);

const fetchCurrencies = async () => {
    try{
      const response = await axios.get<ExchangeRate[]>('/exchangeRates');
      rates.value = response.data;
      currencies.value = [...new Set(response.data.map((rate: ExchangeRate) => [rate.fromcurrency, rate.tocurrency]).flat())];
    }catch(error){
      console.log('Failed to load currencies', error)
    }
};

const exchange = () => {
    const rate = rates.value.find(
        (rate) => rate.fromcurrency === form.value.fromcurrency && rate.tocurrency === form.value.tocurrency
    )?.rate;

    if (rate) {
        result.value = form.value.amount * rate;
    } else {
        result.value = null;
    }
}

onMounted(fetchCurrencies);
</script>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    padding: 20px;
}

.exchange-form {  
    width: 100%;  
    max-width: 600px;  
    margin: 20px auto;  
    padding: 20px;  
    background-color: #f5f5f5;  
    border-radius: 4px;  
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);  
}  

.result { 
    width: 100%; 
    max-width: 600px; 
    padding: 20px;  
    background-color: #f0f0f0;  
    border-radius: 4px;  
    text-align: left;  
    font-size: 18px;  
}  
</style>