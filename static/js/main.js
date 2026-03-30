const initUI = () => {
  const cards = document.querySelectorAll('.reveal');
  if ('IntersectionObserver' in window && cards.length > 0) {
    const observer = new IntersectionObserver((entries, obs) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.style.animationPlayState = 'running';
          obs.unobserve(entry.target);
        }
      });
    }, { threshold: 0.12 });

    cards.forEach((card) => {
      card.style.animationPlayState = 'paused';
      observer.observe(card);
    });
  }

  const shots = document.querySelectorAll('[data-lightbox-shot] img');
  if (shots.length === 0) return;

  const overlay = document.createElement('div');
  overlay.className = 'image-lightbox';
  overlay.innerHTML = '<button class="image-lightbox-close" type="button" aria-label="Fermer">×</button><img alt="Aperçu agrandi">';
  document.body.appendChild(overlay);

  const lightboxImage = overlay.querySelector('img');
  const closeButton = overlay.querySelector('.image-lightbox-close');

  const close = () => overlay.classList.remove('open');

  const open = (img) => {
    lightboxImage.src = img.src;
    lightboxImage.alt = img.alt;
    overlay.classList.add('open');
  };

  shots.forEach((img) => {
    const shot = img.closest('[data-lightbox-shot]');
    if (shot) shot.addEventListener('click', () => open(img));
    img.addEventListener('click', (event) => {
      event.preventDefault();
      open(img);
    });
  });

  document.addEventListener('click', (event) => {
    const targetShot = event.target.closest?.('[data-lightbox-shot] img');
    if (targetShot) {
      event.preventDefault();
      open(targetShot);
    }
  });

  overlay.addEventListener('click', (event) => {
    if (event.target === overlay) close();
  });

  closeButton.addEventListener('click', close);

  document.addEventListener('keydown', (event) => {
    if (event.key === 'Escape') close();
  });
};

if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', initUI);
} else {
  initUI();
}
