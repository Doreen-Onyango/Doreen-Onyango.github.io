document.addEventListener("DOMContentLoaded", (event) => {
  // Smooth scrolling for navigation links
  document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
    anchor.addEventListener("click", function (e) {
      e.preventDefault();
      document.querySelector(this.getAttribute("href")).scrollIntoView({
        behavior: "smooth",
      });
    });
  });

  // Animate skill bars on scroll
  const skillBars = document.querySelectorAll(".skill-progress");
  const animateSkillBars = () => {
    skillBars.forEach((bar) => {
      const barTop = bar.getBoundingClientRect().top;
      if (barTop < window.innerHeight - 50) {
        bar.style.width = bar.parentElement.getAttribute("data-progress") + "%";
      }
    });
  };

  window.addEventListener("scroll", animateSkillBars);
  animateSkillBars(); // Initial check on page load

  // Form submission handling
  const contactForm = document.querySelector(".contact-form");
  contactForm.addEventListener("submit", (e) => {
    // e.preventDefault();
    // Here you would typically send the form data to a server
    // alert("Thank you for your message! I'll get back to you soon.");
    // contactForm.reset();
  });

  // Easter egg: Konami code
  let konamiCode = [
    "ArrowUp",
    "ArrowUp",
    "ArrowDown",
    "ArrowDown",
    "ArrowLeft",
    "ArrowRight",
    "ArrowLeft",
    "ArrowRight",
    "b",
    "a",
  ];
  let konamiCodePosition = 0;

  document.addEventListener("keydown", (e) => {
    if (e.key === konamiCode[konamiCodePosition]) {
      konamiCodePosition++;
      if (konamiCodePosition === konamiCode.length) {
        activateEasterEgg();
        konamiCodePosition = 0;
      }
    } else {
      konamiCodePosition = 0;
    }
  });

  function activateEasterEgg() {
    document.body.style.fontFamily = '"Comic Sans MS", cursive';
    alert(
      "Congratulations! You've unlocked the secret font mode. Enjoy the nostalgia!"
    );
    setTimeout(() => {
      document.body.style.fontFamily = "";
    }, 5000);
  }

  function animateBackgroundCurves() {
    const curves = document.querySelector(".background-curves svg path");
    curves.animate(
      [
        {
          d: "M0,96L48,112C96,128,192,160,288,186.7C384,213,480,235,576,213.3C672,192,768,128,864,101.3C960,75,1056,85,1152,101.3C1248,117,1344,139,1392,149.3L1440,160L1440,320L1392,320C1344,320,1248,320,1152,320C1056,320,960,320,864,320C768,320,672,320,576,320C480,320,384,320,288,320C192,320,96,320,48,320L0,320Z",
        },
        {
          d: "M0,160L48,138.7C96,117,192,75,288,64C384,53,480,75,576,101.3C672,128,768,160,864,165.3C960,171,1056,149,1152,144C1248,139,1344,149,1392,154.7L1440,160L1440,320L1392,320C1344,320,1248,320,1152,320C1056,320,960,320,864,320C768,320,672,320,576,320C480,320,384,320,288,320C192,320,96,320,48,320L0,320Z",
        },
        {
          d: "M0,96L48,112C96,128,192,160,288,186.7C384,213,480,235,576,213.3C672,192,768,128,864,101.3C960,75,1056,85,1152,101.3C1248,117,1344,139,1392,149.3L1440,160L1440,320L1392,320C1344,320,1248,320,1152,320C1056,320,960,320,864,320C768,320,672,320,576,320C480,320,384,320,288,320C192,320,96,320,48,320L0,320Z",
        },
      ],
      {
        duration: 20000,
        iterations: Infinity,
        easing: "ease-in-out",
      }
    );
  }

  animateBackgroundCurves();

  // Add hover effects to project cards
  const projectCards = document.querySelectorAll(".project-card");
  projectCards.forEach((card) => {
    card.addEventListener("mouseenter", () => {
      card.style.transform = "translateY(-10px) rotate(2deg)";
    });
    card.addEventListener("mouseleave", () => {
      card.style.transform = "translateY(0) rotate(0)";
    });
  });
});

// drip any when page is opened
document.getElementById("contact-info").addEventListener("click", function() {
  const icons = document.querySelectorAll(".contact-icon");
  
  icons.forEach((icon, index) => {
      setTimeout(() => {
          icon.classList.add("drip-in");
      }, index * 300); // Delay each icon by 300ms
  });
});

function toggleMenu() {
  const navMenu = document.getElementById("nav-menu");
  navMenu.classList.toggle("nav-active");
}

// Close the menu when any nav link is clicked
document.querySelectorAll('#nav-menu a').forEach(link => {
  link.addEventListener('click', () => {
    const navMenu = document.getElementById("nav-menu");
    navMenu.classList.remove("nav-active");
  });
});

const themeToggle = document.getElementById("theme-toggle");

themeToggle.addEventListener("click", () => {
  document.body.classList.toggle("dark-theme");

  // Change the button text based on the theme
  if (document.body.classList.contains("dark-theme")) {
    themeToggle.textContent = "☀️ ";
  } else {
    themeToggle.textContent = "🌙";
  }
});
