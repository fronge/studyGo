"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.formatTime = void 0;
const formatTime = (date) => {
    const hour = date.getHours();
    const minute = date.getMinutes();
    const second = date.getSeconds();
    return ([hour, minute, second].map(formatNumber).join(':'));
};
exports.formatTime = formatTime;
const formatNumber = (n) => {
    const s = n.toString();
    return s[1] ? s : '0' + s;
};
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidXRpbC5qcyIsInNvdXJjZVJvb3QiOiIiLCJzb3VyY2VzIjpbInV0aWwudHMiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6Ijs7O0FBQU8sTUFBTSxVQUFVLEdBQUcsQ0FBQyxJQUFVLEVBQUUsRUFBRTtJQUl2QyxNQUFNLElBQUksR0FBRyxJQUFJLENBQUMsUUFBUSxFQUFFLENBQUE7SUFDNUIsTUFBTSxNQUFNLEdBQUcsSUFBSSxDQUFDLFVBQVUsRUFBRSxDQUFBO0lBQ2hDLE1BQU0sTUFBTSxHQUFHLElBQUksQ0FBQyxVQUFVLEVBQUUsQ0FBQTtJQUVoQyxPQUFPLENBQ0wsQ0FBQyxJQUFJLEVBQUUsTUFBTSxFQUFFLE1BQU0sQ0FBQyxDQUFDLEdBQUcsQ0FBQyxZQUFZLENBQUMsQ0FBQyxJQUFJLENBQUMsR0FBRyxDQUFDLENBQ25ELENBQUE7QUFDSCxDQUFDLENBQUE7QUFYWSxRQUFBLFVBQVUsY0FXdEI7QUFFRCxNQUFNLFlBQVksR0FBRyxDQUFDLENBQVMsRUFBRSxFQUFFO0lBQ2pDLE1BQU0sQ0FBQyxHQUFHLENBQUMsQ0FBQyxRQUFRLEVBQUUsQ0FBQTtJQUN0QixPQUFPLENBQUMsQ0FBQyxDQUFDLENBQUMsQ0FBQyxDQUFDLENBQUMsQ0FBQyxDQUFDLENBQUMsQ0FBQyxHQUFHLEdBQUcsQ0FBQyxDQUFBO0FBQzNCLENBQUMsQ0FBQSIsInNvdXJjZXNDb250ZW50IjpbImV4cG9ydCBjb25zdCBmb3JtYXRUaW1lID0gKGRhdGU6IERhdGUpID0+IHtcbiAgLy8gY29uc3QgeWVhciA9IGRhdGUuZ2V0RnVsbFllYXIoKVxuICAvLyBjb25zdCBtb250aCA9IGRhdGUuZ2V0TW9udGgoKSArIDFcbiAgLy8gY29uc3QgZGF5ID0gZGF0ZS5nZXREYXRlKClcbiAgY29uc3QgaG91ciA9IGRhdGUuZ2V0SG91cnMoKVxuICBjb25zdCBtaW51dGUgPSBkYXRlLmdldE1pbnV0ZXMoKVxuICBjb25zdCBzZWNvbmQgPSBkYXRlLmdldFNlY29uZHMoKVxuXG4gIHJldHVybiAoXG4gICAgW2hvdXIsIG1pbnV0ZSwgc2Vjb25kXS5tYXAoZm9ybWF0TnVtYmVyKS5qb2luKCc6JylcbiAgKVxufVxuXG5jb25zdCBmb3JtYXROdW1iZXIgPSAobjogbnVtYmVyKSA9PiB7XG4gIGNvbnN0IHMgPSBuLnRvU3RyaW5nKClcbiAgcmV0dXJuIHNbMV0gPyBzIDogJzAnICsgc1xufVxuXG4iXX0=