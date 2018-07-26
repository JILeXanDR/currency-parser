(function () {

    let initWebSocketConnection = function connect(onmessage) {
        let socket = new WebSocket("ws://localhost:12345/io");

        socket.onopen = function () {
            console.info('connected');
        };

        socket.onclose = function (e) {
            console.error(`connection closed (${e.code})`);
            setTimeout(() => connect(onmessage), 1000)
        };

        socket.onmessage = onmessage;
    };

    Vue.component('chart', {
        extends: VueChartJs.Line,
        mixins: [VueChartJs.mixins.reactiveProp],
        props: ['options'],
        mounted() {
            this.renderChart(this.chartData, this.options)
        },
        watch: {
            'options': {
                handler(newOptions) {
                    this.renderChart(this.chartData, newOptions)
                },
                deep: true,
            }
        },
    });

    var App = Vue.component('app', {
        template: '#tpl_app',
        data() {
            return {
                chart: {
                    data: {
                        showLine: true,
                        datasets: [
                            {
                                label: 'мин.',
                                borderColor: '#0cf811',
                                borderWidth: 1,
                                showLine: true, // show line in scatter plot
                                fill: false, // only show line
                                data: [],
                            },
                            {
                                label: 'сред.',
                                borderColor: '#f83b60',
                                borderWidth: 1,
                                showLine: true, // show line in scatter plot
                                fill: false, // only show line
                                data: [],
                            },
                            {
                                label: 'макс.',
                                borderColor: '#0032ff',
                                borderWidth: 1,
                                showLine: true, // show line in scatter plot
                                fill: false, // only show line
                                data: [],
                            },
                        ],
                    },
                    options: {
                        animation: false,
                        responsive: true,
                        maintainAspectRatio: false,
                        scales: {
                            xAxes: [{
                                type: "time",
                                time: {
                                    unit: 'second',
                                    unitStepSize: 1,
                                    round: 'second',
                                    tooltipFormat: "hh:mm:ss",
                                    displayFormats: {
                                        hour: 'hh:mm'
                                    },
                                    min: moment(),
                                    max: moment(),
                                }
                            }]
                        }
                    },
                },
            }
        },
        created() {

            let updateMinMaxTime = (maxTime) => {
                let now = moment(maxTime);
                let temp = this.chart.options;
                temp.scales.xAxes[0].time.min = now.clone().add(-30, 'seconds');
                temp.scales.xAxes[0].time.max = now.clone();

                this.chart.options = temp;
            };

            updateMinMaxTime();

            initWebSocketConnection((e) => {
                    const {event, payload} = JSON.parse(e.data);

                    if (event === 'rates') {

                        payload.forEach((data) => {

                            let f = function (rate) {
                                return {
                                    t: moment(data.time), // time
                                    y: data[rate], // rate
                                };
                            };

                            let min = this.chart.data.datasets[0];
                            let avg = this.chart.data.datasets[1];
                            let max = this.chart.data.datasets[2];
                            this.chart.data = {
                                ...this.chart.data,
                                datasets: [
                                    {...min, data: min.data.concat([f('min_rate')])},
                                    {...avg, data: avg.data.concat([f('avg_rate')])},
                                    {...max, data: max.data.concat([f('max_rate')])},
                                ],
                            };

                            updateMinMaxTime(data.time);
                        });
                    }
                }
            );
        },
    });

    new Vue({
        el: '#app',
        render: (h) => h(App)
    });
})();
