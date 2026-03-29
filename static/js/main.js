document.addEventListener('DOMContentLoaded', () => {
  const cards = document.querySelectorAll('.reveal');
  if (!('IntersectionObserver' in window) || cards.length === 0) return;

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
});
