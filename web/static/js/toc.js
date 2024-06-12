const header = document.getElementById('table-of-contents');
const toc = document.getElementById('toc-list');

header.addEventListener('click', toggle);

function toggle() {
  if (toc.classList.contains('expanded')) {
    header.classList.remove('expanded');
    toc.classList.remove('expanded');
  } else {
    header.classList.add('expanded');
    toc.classList.add('expanded');
  }
}
