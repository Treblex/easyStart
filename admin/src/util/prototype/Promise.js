// 无论如何执行 
Promise.prototype.finally = function (callback) {
    let P = this.constructor;
    return this.then(
        value => P.resolve(callback()).then(() => value),
        reason => {
            P.resolve(callback()).then(() => {
                throw reason;
            })
        }
    );
};

// 使Promise结束后抛出全局错误
Promise.prototype.done = function (onFulfilled, onRejected) {
    this.then(onFulfilled, onRejected)
        .catch(function (reason) {
            // 抛出一个全局错误
            setTimeout(() => { throw reason }, 0);
        });
};