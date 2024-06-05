init();

function init() {
  const urlParams = new URLSearchParams(location.search);
  document.querySelector('input[name="search"]').value = urlParams.get('search') || '';
  document.querySelector('select[name="catalog"]').value = urlParams.get('catalog') || '';

  document.querySelector('input[name="search"]').addEventListener('keyup', update);
  document.querySelector('input[name="search"]').addEventListener('change', update);
  document.querySelector('select[name="catalog"]').addEventListener('change', update);
}

function update() {
  const filters = {
    search: document.querySelector('input[name="search"]').value,
    catalog: document.querySelector('select[name="catalog"]').value,
  };
  const cleanFilters = Object.fromEntries(Object.entries(filters).filter(([_, v]) => v != null && v !== ''));
  const queryParams = new URLSearchParams(cleanFilters).toString();

  history.pushState(null, '', queryParams ? '/books?' + queryParams : '/books');
  debouncedFetchBooks(queryParams);
}

const debouncedFetchBooks = debounce(fetchBooks, 500);

function fetchBooks(queryParams) {
  fetch('_books?' + queryParams)
    .then(response => response.text())
    .then(html => document.querySelector('#book-list').innerHTML = html)
    .catch(error => console.error('Error:', error));
}

function debounce(func, delay) {
  let timeoutId;
  return function (...args) {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
      func.apply(this, args);
    }, delay);
  };
}
