async function fetchCurrencies() {
    const url = 'https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1';
    try {
        const response = await fetch(url);
        const data = await response.json();
        displayCurrencies(data);
    } catch (error) {
        console.error('Ошибка при получении данных:', error);
    }
}

function displayCurrencies(data) {
    const table = document.getElementById('currency-table');
    const currencyList = document.getElementById('currency-list');

    data.slice(0, 5).forEach(currency => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${currency.id}</td>
            <td>${currency.symbol}</td>
            <td>${currency.name}</td>
        `;
        row.classList.add('blue-background');
        currencyList.appendChild(row);
    });

    data.forEach(currency => {
        if (currency.symbol === 'usdt') {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${currency.id}</td>
                <td>${currency.symbol}</td>
                <td>${currency.name}</td>
            `;
            row.classList.add('green-background');
            currencyList.appendChild(row);
        }
    });

    table.style.display = 'block';
}

fetchCurrencies();
