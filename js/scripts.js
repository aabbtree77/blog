document.addEventListener('DOMContentLoaded', () => {
    document.querySelectorAll('.read-more').forEach(button => {
        button.addEventListener('click', () => {
            const article = button.closest('.article');
            const summary = article.querySelector('.summary');
            const content = article.querySelector('.content');
            if (content.style.display === 'none' || content.style.display === '') {
                content.style.display = 'block';
                summary.style.display = 'none';
                button.textContent = 'Read less';
            } else {
                content.style.display = 'none';
                summary.style.display = 'block';
                button.textContent = 'Read more';
            }
        });
    });

    document.querySelectorAll('.content').forEach(content => {
        content.style.display = 'none';
    });
});

