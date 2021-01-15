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
    EntUsertypeEdges,
    EntUsertypeEdgesFromJSON,
    EntUsertypeEdgesFromJSONTyped,
    EntUsertypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUsertype
 */
export interface EntUsertype {
    /**
     * 
     * @type {EntUsertypeEdges}
     * @memberof EntUsertype
     */
    edges?: EntUsertypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntUsertype
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntUsertype
     */
    name?: string;
}

export function EntUsertypeFromJSON(json: any): EntUsertype {
    return EntUsertypeFromJSONTyped(json, false);
}

export function EntUsertypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUsertype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntUsertypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function EntUsertypeToJSON(value?: EntUsertype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntUsertypeEdgesToJSON(value.edges),
        'id': value.id,
        'name': value.name,
    };
}

