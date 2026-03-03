// Smooth scroll for navigation links
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            target.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            });
        }
    });
});

// Intersection Observer for fade-in animations
const observerOptions = {
    threshold: 0.1,
    rootMargin: '0px 0px -50px 0px'
};

const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.classList.add('visible');

            // Animate feature cards sequentially
            if (entry.target.classList.contains('features-grid')) {
                const cards = entry.target.querySelectorAll('.feature-card');
                cards.forEach((card, index) => {
                    setTimeout(() => {
                        card.classList.add('animate-in');
                    }, index * 100);
                });
            }

            // Animate steps sequentially
            if (entry.target.classList.contains('steps-container')) {
                const steps = entry.target.querySelectorAll('.step');
                steps.forEach((step, index) => {
                    setTimeout(() => {
                        step.classList.add('animate-in');
                    }, index * 200);
                });
            }
        }
    });
}, observerOptions);

// Observe sections for animations
document.querySelectorAll('.features, .how-it-works, .cta-section').forEach(section => {
    observer.observe(section);
});

// Parallax effect for gradient orbs
let mouseX = 0;
let mouseY = 0;
let targetX = 0;
let targetY = 0;

document.addEventListener('mousemove', (e) => {
    mouseX = (e.clientX / window.innerWidth - 0.5) * 20;
    mouseY = (e.clientY / window.innerHeight - 0.5) * 20;
});

function animateOrbs() {
    targetX += (mouseX - targetX) * 0.05;
    targetY += (mouseY - targetY) * 0.05;

    const orbs = document.querySelectorAll('.gradient-orb');
    orbs.forEach((orb, index) => {
        const speed = (index + 1) * 0.5;
        orb.style.transform = `translate(${targetX * speed}px, ${targetY * speed}px)`;
    });

    requestAnimationFrame(animateOrbs);
}

animateOrbs();

// Navbar scroll effect
let lastScroll = 0;
const navbar = document.querySelector('.navbar');

window.addEventListener('scroll', () => {
    const currentScroll = window.pageYOffset;

    if (currentScroll <= 0) {
        navbar.style.boxShadow = 'none';
    } else {
        navbar.style.boxShadow = '0 2px 20px rgba(0, 0, 0, 0.5)';
    }

    if (currentScroll > lastScroll && currentScroll > 100) {
        navbar.style.transform = 'translateY(-100%)';
    } else {
        navbar.style.transform = 'translateY(0)';
    }

    lastScroll = currentScroll;
});

// Add hover effect to preview cards
const previewCards = document.querySelectorAll('.preview-card');
previewCards.forEach(card => {
    card.addEventListener('mouseenter', () => {
        card.style.transform = 'scale(1.05)';
    });
    card.addEventListener('mouseleave', () => {
        card.style.transform = 'scale(1)';
    });
});

// Typewriter effect for dynamic text
class TypeWriter {
    constructor(element, texts, speed = 100) {
        this.element = element;
        this.texts = texts;
        this.speed = speed;
        this.textIndex = 0;
        this.charIndex = 0;
        this.isDeleting = false;
        this.init();
    }

    init() {
        this.type();
    }

    type() {
        const currentText = this.texts[this.textIndex];
        const displayText = currentText.substring(0, this.charIndex);

        this.element.textContent = displayText;

        if (!this.isDeleting && this.charIndex === currentText.length) {
            // Pause at the end of the text
            setTimeout(() => {
                this.isDeleting = true;
                this.type();
            }, 2000);
            return;
        }

        if (this.isDeleting && this.charIndex === 0) {
            this.isDeleting = false;
            this.textIndex = (this.textIndex + 1) % this.texts.length;
        }

        const typeSpeed = this.isDeleting ? this.speed / 2 : this.speed;
        this.charIndex += this.isDeleting ? -1 : 1;

        setTimeout(() => this.type(), typeSpeed);
    }
}

// Initialize typewriter for dynamic taglines
const dynamicTexts = [
    'Daily Victories',
    'Consistent Progress',
    'Meaningful Results'
];

// Wait for DOM to be fully loaded
document.addEventListener('DOMContentLoaded', () => {
    const gradientText = document.querySelector('.gradient-text');
    if (gradientText) {
        new TypeWriter(gradientText, dynamicTexts, 100);
    }

    // Add counter animation for stats
    const stats = document.querySelectorAll('.stat-number');
    const animateCounter = (element) => {
        const target = element.textContent;
        const isPercentage = target.includes('%');
        const targetNumber = parseInt(target.replace(/[^0-9]/g, ''));
        const suffix = target.replace(/[0-9]/g, '');
        let current = 0;
        const increment = targetNumber / 100;
        const timer = setInterval(() => {
            current += increment;
            if (current >= targetNumber) {
                current = targetNumber;
                clearInterval(timer);
            }
            element.textContent = Math.floor(current) + suffix;
        }, 10);
    };

    const statsObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting && !entry.target.classList.contains('animated')) {
                animateCounter(entry.target);
                entry.target.classList.add('animated');
            }
        });
    }, { threshold: 0.5 });

    stats.forEach(stat => statsObserver.observe(stat));

    // Progress bar animation
    const progressFill = document.querySelector('.progress-fill');
    if (progressFill) {
        const progressObserver = new IntersectionObserver((entries) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    entry.target.style.animation = 'fillProgress 2s ease-out forwards';
                }
            });
        }, { threshold: 0.5 });

        progressObserver.observe(progressFill);
    }
});

// Add CSS for animations dynamically
const style = document.createElement('style');
style.textContent = `
    @keyframes fillProgress {
        from { width: 0; }
        to { width: 50%; }
    }

    .features-grid .feature-card {
        opacity: 0;
        transform: translateY(20px);
        transition: all 0.6s ease-out;
    }

    .features-grid .feature-card.animate-in {
        opacity: 1;
        transform: translateY(0);
    }

    .steps-container .step {
        opacity: 0;
        transform: translateY(20px);
        transition: all 0.6s ease-out;
    }

    .steps-container .step.animate-in {
        opacity: 1;
        transform: translateY(0);
    }
`;
document.head.appendChild(style);

// CTA button ripple effect
document.querySelectorAll('.btn-gradient, .btn-primary').forEach(button => {
    button.addEventListener('click', function(e) {
        const ripple = document.createElement('span');
        ripple.classList.add('ripple');

        const rect = this.getBoundingClientRect();
        const size = Math.max(rect.width, rect.height);
        const x = e.clientX - rect.left - size / 2;
        const y = e.clientY - rect.top - size / 2;

        ripple.style.width = ripple.style.height = size + 'px';
        ripple.style.left = x + 'px';
        ripple.style.top = y + 'px';

        this.appendChild(ripple);

        setTimeout(() => ripple.remove(), 600);
    });
});

// Add ripple CSS
const rippleStyle = document.createElement('style');
rippleStyle.textContent = `
    .btn-gradient, .btn-primary {
        position: relative;
        overflow: hidden;
    }

    .ripple {
        position: absolute;
        border-radius: 50%;
        background: rgba(255, 255, 255, 0.5);
        transform: scale(0);
        animation: rippleAnimation 0.6s ease-out;
        pointer-events: none;
    }

    @keyframes rippleAnimation {
        to {
            transform: scale(4);
            opacity: 0;
        }
    }
`;
document.head.appendChild(rippleStyle);

// Add loading animation for page load
window.addEventListener('load', () => {
    document.body.classList.add('loaded');
});

// Console Easter egg
console.log(
    '%c🔥 Welcome to Fireplace! 🔥',
    'font-size: 24px; font-weight: bold; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; padding: 10px 20px; border-radius: 10px;'
);
console.log(
    '%cTransform your goals into daily victories',
    'font-size: 16px; color: #8b5cf6; font-style: italic;'
);