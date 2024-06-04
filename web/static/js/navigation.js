const menu = document.getElementById("menu-button");
const logo = document.getElementById("brand-logo");
const container = document.getElementById("menu-container");

menu.addEventListener("click", function () {
  if (container.hidden) {
    showMenu();
  } else {
    hideMenu();
  }
  menu.focus();
});

document.addEventListener("keyup", function (e) {
  if (e.key === "Escape" && !container.hidden) {
    hideMenu()
    menu.focus();
  }
});

const currentPath = window.location.pathname;
const menuItems = document.querySelectorAll("li[data-path]");

menuItems.forEach((item) => {
  const path = item.getAttribute("data-path");
  const pathSegments = path.split('/').filter(segment => segment);
  const currentPathSegments = currentPath.split('/').filter(segment => segment);

  if (pathSegments[0] === currentPathSegments[0]) {
    item.classList.add("text-orange-500");
    item.classList.remove("text-white");
  }
});

function hideMenu() {
  const spanClass = "inline-block rounded-sm h-[4px] bg-white mx-auto my-0.5";
  logo.hidden = false;
  container.hidden = true;
  menu.children[0].className = spanClass + " w-6";
  menu.children[1].className = spanClass + " w-4 group-hover:w-6";
  menu.children[2].className = spanClass + " w-2 group-hover:w-6";
}

function showMenu() {
  const spanClass = "inline-block rounded-sm h-[4px] bg-white mx-auto my-0.5";
  logo.hidden = true;
  container.hidden = false;
  menu.children[0].className = spanClass + " w-6 rotate-45 translate-y-[10px]";
  menu.children[1].className = spanClass + " w-[10px] -rotate-45 translate-x-[-6px] translate-y-2";
  menu.children[2].className = spanClass + " w-[10px] -rotate-45 translate-x-[6px] -translate-y-3";
}
