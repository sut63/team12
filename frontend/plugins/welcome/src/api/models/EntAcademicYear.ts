/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntAcademicYearEdges,
    EntAcademicYearEdgesFromJSON,
    EntAcademicYearEdgesFromJSONTyped,
    EntAcademicYearEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntAcademicYear
 */
export interface EntAcademicYear {
    /**
     * 
     * @type {EntAcademicYearEdges}
     * @memberof EntAcademicYear
     */
    edges?: EntAcademicYearEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntAcademicYear
     */
    id?: number;
    /**
     * Semester holds the value of the "semester" field.
     * @type {string}
     * @memberof EntAcademicYear
     */
    semester?: string;
}

export function EntAcademicYearFromJSON(json: any): EntAcademicYear {
    return EntAcademicYearFromJSONTyped(json, false);
}

export function EntAcademicYearFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntAcademicYear {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntAcademicYearEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'semester': !exists(json, 'semester') ? undefined : json['semester'],
    };
}

export function EntAcademicYearToJSON(value?: EntAcademicYear | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntAcademicYearEdgesToJSON(value.edges),
        'id': value.id,
        'semester': value.semester,
    };
}

