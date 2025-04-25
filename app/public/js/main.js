// 网页加载完成后执行
document.addEventListener('DOMContentLoaded', function() {
    console.log('网页已加载完成');
    
    // 绑定按钮点击事件
    const buttons = document.querySelectorAll('.btn');
    buttons.forEach(function(button) {
        button.addEventListener('click', function(e) {
            console.log('按钮被点击:', this.textContent);
        });
    });
    
    // 显示当前时间
    const updateTime = function() {
        const now = new Date();
        const timeString = now.toLocaleTimeString();
        const dateElement = document.getElementById('current-time');
        
        if (dateElement) {
            dateElement.textContent = timeString;
        }
    };
    
    // 初始更新时间
    updateTime();
    
    // 每秒更新一次时间
    setInterval(updateTime, 1000);
    
    // 简单的表单验证
    const forms = document.querySelectorAll('form');
    forms.forEach(function(form) {
        form.addEventListener('submit', function(e) {
            const requiredFields = form.querySelectorAll('[required]');
            let isValid = true;
            
            requiredFields.forEach(function(field) {
                if (!field.value.trim()) {
                    isValid = false;
                    field.classList.add('error');
                } else {
                    field.classList.remove('error');
                }
            });
            
            if (!isValid) {
                e.preventDefault();
                alert('请填写所有必填字段');
            }
        });
    });
}); 