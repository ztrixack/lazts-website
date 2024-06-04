const theme = document.getElementById('theme-toggle');

function updateTheme() {
  const isDark = document.documentElement.classList.toggle('dark');
  localStorage.setItem('darkMode', isDark); // Store the state in localStorage

  // Update classes and styles for the theme toggle button
  if (isDark) {
    theme.classList.replace('bg-orange-300', 'bg-transparent');
    theme.style.boxShadow = "inset -9px -8px 1px 1px #ddd";
    reloadGiscus('dark');
  } else {
    theme.classList.replace('bg-transparent', 'bg-orange-300');
    theme.style.boxShadow = "";
    reloadGiscus('light');
  }
}

function reloadGiscus(theme) {
  const currentScript = document.querySelector('script[id="comments"]');
  if (!currentScript) return;

  const newScript = document.createElement('script');
  newScript.setAttribute('id', 'comments');
  newScript.setAttribute('src', 'https://giscus.app/client.js');
  newScript.setAttribute('data-repo', 'ztrixack/lazts-giscus');
  newScript.setAttribute('data-repo-id', "R_kgDOL5NElw");
  newScript.setAttribute('data-category', "Comments");
  newScript.setAttribute('data-category-id', "DIC_kwDOL5NEl84CfPg3");
  newScript.setAttribute('data-mapping', "pathname");
  newScript.setAttribute('data-strict', "0");
  newScript.setAttribute('data-reactions-enabled', "1");
  newScript.setAttribute('data-emit-metadata', "0");
  newScript.setAttribute('data-input-position', "bottom");
  newScript.setAttribute('data-theme', theme);
  newScript.setAttribute('data-lang', "th");
  newScript.setAttribute('crossorigin', 'anonymous');
  newScript.async = true;

  currentScript.parentNode.replaceChild(newScript, currentScript);
}

document.addEventListener('DOMContentLoaded', function () {
  theme.addEventListener('click', updateTheme);

  // Set initial theme from local storage
  if (localStorage.getItem('darkMode') === 'false') {
    document.documentElement.classList.add('dark');
    updateTheme();
  } else {
    document.documentElement.classList.remove('dark');
    updateTheme();
  }
})
