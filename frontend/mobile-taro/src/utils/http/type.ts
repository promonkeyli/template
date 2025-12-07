export interface ResponseData<T = any> {
    code: number;        // 业务状态码
    message: string;     // 描述信息
    data: T;             // 实际数据
}
