now = moment

moment.prototype.sub = function (duration) {
    return moment(this).subtract(moment.duration(duration));
}
