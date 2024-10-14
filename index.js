// Scroll Animations
function isElementInViewport(el) {
    const rect = el.getBoundingClientRect();
    return (
      rect.top >= 0 &&
      rect.left >= 0 &&
      rect.bottom <=
        (window.innerHeight || document.documentElement.clientHeight) &&
      rect.right <= (window.innerWidth || document.documentElement.clientWidth)
    );
  }
  
  function handleScroll() {
    const sections = document.querySelectorAll("section");
    sections.forEach((section) => {
      if (isElementInViewport(section)) {
        section.classList.add("visible");
      }
    });
  
    const projectCards = document.querySelectorAll(".project-card");
    projectCards.forEach((card, index) => {
      if (isElementInViewport(card)) {
        setTimeout(() => {
          card.classList.add("visible");
        }, index * 200);
      }
    });
  }
  
  window.addEventListener("scroll", handleScroll);
  window.addEventListener("load", handleScroll);
  
  // Form Submission
  const form = document.getElementById("contact-form");
  form.addEventListener("submit", function (e) {
    e.preventDefault();
    alert("Thank you for your message! I'll get back to you soon.");
    form.reset();
  });
  
  // Lightbox for Creative Highlights
  const galleryItems = document.querySelectorAll(".gallery-item");
  galleryItems.forEach((item) => {
    item.addEventListener("click", function () {
      const img = this.querySelector("img");
      const lightbox = document.createElement("div");
      lightbox.id = "lightbox";
      document.body.appendChild(lightbox);
  
      const lightboxImage = document.createElement("img");
      lightboxImage.src = img.src;
      lightbox.appendChild(lightboxImage);
  
      lightbox.addEventListener("click", (e) => {
        if (e.target !== e.currentTarget) return;
        lightbox.remove();
      });
    });
  });
  
  // Smooth Scroll to Landing Section
  document.addEventListener("DOMContentLoaded", function () {
    const aboutBtn = document.getElementById("about-btn");
    const landingSection = document.getElementById("landing");
  
    aboutBtn.addEventListener("click", function () {
      landingSection.scrollIntoView({ behavior: "smooth" });
    });
  });
  