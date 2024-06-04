document.addEventListener('DOMContentLoaded', function() {
  let prevScrollY = 0;
  let isHidden = false;

  function handleScroll() {
    const currentScrollY = window.scrollY;

    if (currentScrollY > prevScrollY) {
      if (!isHidden) hideDiv();
    } else {
      if (isHidden) showDiv();
    }

    prevScrollY = currentScrollY;
  }

  function hideDiv() {
    const div = document.getElementById('header');
    if (div) {
      div.style.transform = 'translateY(-100%)';
      isHidden = true;
    }
  }

  function showDiv() {
    const div = document.getElementById('header');
    if (div) {
      div.style.transform = 'translateY(0)';
      isHidden = false;
    }
  }

  document.addEventListener('scroll', handleScroll);

  return () => {
    document.removeEventListener('scroll', handleScroll);
  };
});
