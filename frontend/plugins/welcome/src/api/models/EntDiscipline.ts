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
    EntDisciplineEdges,
    EntDisciplineEdgesFromJSON,
    EntDisciplineEdgesFromJSONTyped,
    EntDisciplineEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDiscipline
 */
export interface EntDiscipline {
    /**
     * Discipline holds the value of the "discipline" field.
     * @type {string}
     * @memberof EntDiscipline
     */
    discipline?: string;
    /**
     * 
     * @type {EntDisciplineEdges}
     * @memberof EntDiscipline
     */
    edges?: EntDisciplineEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDiscipline
     */
    id?: number;
}

export function EntDisciplineFromJSON(json: any): EntDiscipline {
    return EntDisciplineFromJSONTyped(json, false);
}

export function EntDisciplineFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDiscipline {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'discipline': !exists(json, 'discipline') ? undefined : json['discipline'],
        'edges': !exists(json, 'edges') ? undefined : EntDisciplineEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDisciplineToJSON(value?: EntDiscipline | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'discipline': value.discipline,
        'edges': EntDisciplineEdgesToJSON(value.edges),
        'id': value.id,
    };
}

